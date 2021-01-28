package robot

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	"github.com/edaniels/gostream/codec/vpx"

	"github.com/echolabsinc/robotcore/base"
	"github.com/echolabsinc/robotcore/lidar"
	"github.com/echolabsinc/robotcore/utils/stream"
)

type robotWebApp struct {
	template    *template.Template
	remoteViews []gostream.RemoteView
	theRobot    *Robot
}

func (app *robotWebApp) Init() error {
	_, thisFilePath, _, _ := runtime.Caller(0)
	thisDirPath, err := filepath.Abs(filepath.Dir(thisFilePath))
	if err != nil {
		return err
	}
	t, err := template.New("foo").Funcs(template.FuncMap{
		"jsSafe": func(js string) template.JS {
			return template.JS(js)
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).ParseGlob(fmt.Sprintf("%s/*.html", thisDirPath))
	if err != nil {
		return err
	}
	app.template = t.Lookup("webappindex.html")
	return nil
}

func (app *robotWebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if true {
		err := app.Init()
		if err != nil {
			golog.Global.Debugf("couldn't reload template: %s", err)
			return
		}
	}

	type RemoteView struct {
		JavaScript string
		Body       string
	}

	type Temp struct {
		RemoteViews []RemoteView
		Bases       []string
		Arms        []string
		Grippers    []string
	}

	temp := Temp{}

	for idx := range app.theRobot.Bases {
		temp.Bases = append(temp.Bases, fmt.Sprintf("base%d", idx))
	}

	for idx := range app.theRobot.Arms {
		temp.Arms = append(temp.Arms, fmt.Sprintf("arm%d", idx))
	}

	for idx := range app.theRobot.Grippers {
		temp.Grippers = append(temp.Grippers, fmt.Sprintf("gripper%d", idx))
	}

	for _, remoteView := range app.remoteViews {
		htmlData := remoteView.HTML()
		temp.RemoteViews = append(temp.RemoteViews, RemoteView{
			htmlData.JavaScript,
			htmlData.Body,
		})
	}

	err := app.template.Execute(w, temp)
	if err != nil {
		golog.Global.Debugf("couldn't execute web page: %s", err)
	}
}

// ---------------

func InstallWebBase(mux *http.ServeMux, theBase base.Base) {

	mux.HandleFunc("/api/base", func(w http.ResponseWriter, req *http.Request) {
		speed := 64 // TODO(erh): this is proably the wrong default
		if req.FormValue("speed") != "" {
			speed2, err := strconv.ParseInt(req.FormValue("speed"), 10, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("bad speed [%s] %s", req.FormValue("speed"), err), http.StatusBadRequest)
				return
			}
			speed = int(speed2)
		}

		s := req.FormValue("stop")
		d := req.FormValue("distanceMM")
		a := req.FormValue("angle")

		var err error

		if s == "t" || s == "true" {
			err = theBase.Stop()
		} else if d != "" {
			d2, err2 := strconv.ParseInt(d, 10, 64)
			if err2 != nil {
				http.Error(w, fmt.Sprintf("bad distance [%s] %s", d, err2), http.StatusBadRequest)
				return
			}

			err = theBase.MoveStraight(int(d2), speed, false)
		} else if a != "" {
			a2, err2 := strconv.ParseInt(a, 10, 64)
			if err2 != nil {
				http.Error(w, fmt.Sprintf("bad angle [%s] %s", d, err2), http.StatusBadRequest)
				return
			}

			err = theBase.Spin(int(a2), speed, false)
		} else {
			http.Error(w, "no stop, distanceMM, angle given", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, fmt.Sprintf("erorr moving %s", err), http.StatusInternalServerError)
		} else {
			_, err = io.WriteString(w, "ok")
			if err != nil {
				panic("impossible")
			}
		}

	})
}

func InstallWebArms(mux *http.ServeMux, theRobot *Robot) {
	mux.HandleFunc("/api/arm", func(w http.ResponseWriter, req *http.Request) {
		var err error

		mode := req.FormValue("mode")
		arm := 0

		if req.FormValue("num") != "" {
			arm2, err := strconv.ParseInt(req.FormValue("num"), 10, 64)
			if err != nil {
				http.Error(w, "bad value for arm", http.StatusBadRequest)
				return
			}
			arm = int(arm2)
		}
		if arm < 0 || arm >= len(theRobot.Arms) {
			http.Error(w, "not a valid arm number", http.StatusBadRequest)
			return
		}

		where := theRobot.Arms[arm].State().CartesianInfo

		if mode == "abs" {
			vals := []int64{}
			for _, n := range []string{"x", "y", "z"} {
				val, err := strconv.ParseInt(req.FormValue(n), 10, 64)
				if err != nil {
					http.Error(w, fmt.Sprintf("bad value for:%s [%s]", n, req.FormValue(n)), http.StatusBadRequest)
					return
				}
				vals = append(vals, val)
			}

			where.X = float64(vals[0]) / 1000
			where.Y = float64(vals[1]) / 1000
			where.Z = float64(vals[2]) / 1000

			err = theRobot.Arms[arm].MoveToPositionC(where)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if mode == "inc" {
			vals := []int64{0, 0, 0}
			for idx, n := range []string{"x", "y", "z"} {
				val, err := strconv.ParseInt(req.FormValue(n), 10, 64)
				if err == nil {
					vals[idx] = val
				}
			}

			where.X += float64(vals[0]) / 1000
			where.Y += float64(vals[1]) / 1000
			where.Z += float64(vals[2]) / 1000

			err = theRobot.Arms[arm].MoveToPositionC(where)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

		clean := map[string]int64{
			"x": int64(where.X * 1000),
			"y": int64(where.Y * 1000),
			"z": int64(where.Z * 1000),
		}

		js, err := json.Marshal(clean)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func InstallWebGrippers(mux *http.ServeMux, theRobot *Robot) {
	mux.HandleFunc("/api/gripper", func(w http.ResponseWriter, req *http.Request) {
		gripper := 0

		if req.FormValue("num") != "" {
			g2, err := strconv.ParseInt(req.FormValue("num"), 10, 64)
			if err != nil {
				http.Error(w, "bad value for num", http.StatusBadRequest)
				return
			}
			gripper = int(g2)
		}

		if gripper < 0 || gripper >= len(theRobot.Grippers) {
			http.Error(w, "not a valid gripper number", http.StatusBadRequest)
			return
		}

		var err error

		action := req.FormValue("action")
		switch action {
		case "open":
			err = theRobot.Grippers[gripper].Open()
		case "grab":
			_, err = theRobot.Grippers[gripper].Grab()
		default:
			err = fmt.Errorf("bad action: (%s)", action)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})
}

// ---------------

func InstallWeb(mux *http.ServeMux, theRobot *Robot) (func(), error) {
	if len(theRobot.Bases) > 1 {
		return nil, fmt.Errorf("robot.InstallWeb robot can't have morem than 1 base right now")
	}

	views := []gostream.RemoteView{}

	// set up camera/lidar streams
	for i := 0; i < len(theRobot.Cameras)+len(theRobot.LidarDevices); i++ {
		config := vpx.DefaultRemoteViewConfig
		config.Debug = false
		config.StreamNumber = i
		var name string
		if i < len(theRobot.Cameras) {
			name = fmt.Sprintf("Camera %d", i+1)
		} else {
			name = fmt.Sprintf("LIDAR %d", i-len(theRobot.Cameras)+1)
		}
		config.StreamName = name
		remoteView, err := gostream.NewRemoteView(config)
		if err != nil {
			return nil, err
		}
		remoteView.SetOnClickHandler(func(x, y int) {
			golog.Global.Debugw("click", "x", x, "y", y)
		})

		views = append(views, remoteView)
	}

	app := &robotWebApp{remoteViews: views, theRobot: theRobot}
	err := app.Init()
	if err != nil {
		return nil, err
	}

	// install routes
	if len(theRobot.Bases) > 0 {
		InstallWebBase(mux, theRobot.Bases[0])
	}

	InstallWebArms(mux, theRobot)

	InstallWebGrippers(mux, theRobot)

	mux.Handle("/", app)

	for _, view := range views {
		handler := view.Handler()
		mux.Handle("/"+handler.Name, handler.Func)
	}

	// start background threads

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	for idx, remoteView := range views {
		if idx < len(theRobot.Cameras) {
			go stream.MatSource(cancelCtx, theRobot.Cameras[idx], remoteView, 33*time.Millisecond, golog.Global)
			continue
		}
		lidarIdx := idx - len(theRobot.Cameras)
		go stream.MatSource(cancelCtx, lidar.NewMatSource(theRobot.LidarDevices[lidarIdx]), remoteView, 33*time.Millisecond, golog.Global)
	}

	return func() {
		cancelFunc()
		for _, v := range views {
			v.Stop()
		}
	}, nil

}

// ---

/*
helper if you don't need to customize at all
*/
func RunWeb(theRobot *Robot) error {
	mux := http.NewServeMux()

	webCloser, err := InstallWeb(mux, theRobot)
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		theRobot.Close()
		webCloser()
		httpServer.Shutdown(context.Background())
	}()

	golog.Global.Debug("going to listen")
	golog.Global.Fatal(httpServer.ListenAndServe())

	return nil
}
