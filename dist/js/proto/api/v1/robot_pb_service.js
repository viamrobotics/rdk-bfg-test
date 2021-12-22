// package: proto.api.v1
// file: proto/api/v1/robot.proto

var proto_api_v1_robot_pb = require("../../../proto/api/v1/robot_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var RobotService = (function () {
  function RobotService() {}
  RobotService.serviceName = "proto.api.v1.RobotService";
  return RobotService;
}());

RobotService.Status = {
  methodName: "Status",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.StatusRequest,
  responseType: proto_api_v1_robot_pb.StatusResponse
};

RobotService.StatusStream = {
  methodName: "StatusStream",
  service: RobotService,
  requestStream: false,
  responseStream: true,
  requestType: proto_api_v1_robot_pb.StatusStreamRequest,
  responseType: proto_api_v1_robot_pb.StatusStreamResponse
};

RobotService.Config = {
  methodName: "Config",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ConfigRequest,
  responseType: proto_api_v1_robot_pb.ConfigResponse
};

RobotService.DoAction = {
  methodName: "DoAction",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.DoActionRequest,
  responseType: proto_api_v1_robot_pb.DoActionResponse
};

RobotService.BaseMoveStraight = {
  methodName: "BaseMoveStraight",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BaseMoveStraightRequest,
  responseType: proto_api_v1_robot_pb.BaseMoveStraightResponse
};

RobotService.BaseMoveArc = {
  methodName: "BaseMoveArc",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BaseMoveArcRequest,
  responseType: proto_api_v1_robot_pb.BaseMoveArcResponse
};

RobotService.BaseSpin = {
  methodName: "BaseSpin",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BaseSpinRequest,
  responseType: proto_api_v1_robot_pb.BaseSpinResponse
};

RobotService.BaseStop = {
  methodName: "BaseStop",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BaseStopRequest,
  responseType: proto_api_v1_robot_pb.BaseStopResponse
};

RobotService.BaseWidthMillis = {
  methodName: "BaseWidthMillis",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BaseWidthMillisRequest,
  responseType: proto_api_v1_robot_pb.BaseWidthMillisResponse
};

RobotService.BoardStatus = {
  methodName: "BoardStatus",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardStatusRequest,
  responseType: proto_api_v1_robot_pb.BoardStatusResponse
};

RobotService.BoardGPIOSet = {
  methodName: "BoardGPIOSet",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardGPIOSetRequest,
  responseType: proto_api_v1_robot_pb.BoardGPIOSetResponse
};

RobotService.BoardGPIOGet = {
  methodName: "BoardGPIOGet",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardGPIOGetRequest,
  responseType: proto_api_v1_robot_pb.BoardGPIOGetResponse
};

RobotService.BoardPWMSet = {
  methodName: "BoardPWMSet",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardPWMSetRequest,
  responseType: proto_api_v1_robot_pb.BoardPWMSetResponse
};

RobotService.BoardPWMSetFrequency = {
  methodName: "BoardPWMSetFrequency",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardPWMSetFrequencyRequest,
  responseType: proto_api_v1_robot_pb.BoardPWMSetFrequencyResponse
};

RobotService.BoardAnalogReaderRead = {
  methodName: "BoardAnalogReaderRead",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardAnalogReaderReadRequest,
  responseType: proto_api_v1_robot_pb.BoardAnalogReaderReadResponse
};

RobotService.BoardDigitalInterruptConfig = {
  methodName: "BoardDigitalInterruptConfig",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardDigitalInterruptConfigRequest,
  responseType: proto_api_v1_robot_pb.BoardDigitalInterruptConfigResponse
};

RobotService.BoardDigitalInterruptValue = {
  methodName: "BoardDigitalInterruptValue",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardDigitalInterruptValueRequest,
  responseType: proto_api_v1_robot_pb.BoardDigitalInterruptValueResponse
};

RobotService.BoardDigitalInterruptTick = {
  methodName: "BoardDigitalInterruptTick",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.BoardDigitalInterruptTickRequest,
  responseType: proto_api_v1_robot_pb.BoardDigitalInterruptTickResponse
};

RobotService.SensorReadings = {
  methodName: "SensorReadings",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.SensorReadingsRequest,
  responseType: proto_api_v1_robot_pb.SensorReadingsResponse
};

RobotService.CompassHeading = {
  methodName: "CompassHeading",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.CompassHeadingRequest,
  responseType: proto_api_v1_robot_pb.CompassHeadingResponse
};

RobotService.CompassStartCalibration = {
  methodName: "CompassStartCalibration",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.CompassStartCalibrationRequest,
  responseType: proto_api_v1_robot_pb.CompassStartCalibrationResponse
};

RobotService.CompassStopCalibration = {
  methodName: "CompassStopCalibration",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.CompassStopCalibrationRequest,
  responseType: proto_api_v1_robot_pb.CompassStopCalibrationResponse
};

RobotService.CompassMark = {
  methodName: "CompassMark",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.CompassMarkRequest,
  responseType: proto_api_v1_robot_pb.CompassMarkResponse
};

RobotService.ForceMatrixMatrix = {
  methodName: "ForceMatrixMatrix",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ForceMatrixMatrixRequest,
  responseType: proto_api_v1_robot_pb.ForceMatrixMatrixResponse
};

RobotService.ForceMatrixSlipDetection = {
  methodName: "ForceMatrixSlipDetection",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ForceMatrixSlipDetectionRequest,
  responseType: proto_api_v1_robot_pb.ForceMatrixSlipDetectionResponse
};

RobotService.ExecuteFunction = {
  methodName: "ExecuteFunction",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ExecuteFunctionRequest,
  responseType: proto_api_v1_robot_pb.ExecuteFunctionResponse
};

RobotService.ExecuteSource = {
  methodName: "ExecuteSource",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ExecuteSourceRequest,
  responseType: proto_api_v1_robot_pb.ExecuteSourceResponse
};

RobotService.InputControllerControls = {
  methodName: "InputControllerControls",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.InputControllerControlsRequest,
  responseType: proto_api_v1_robot_pb.InputControllerControlsResponse
};

RobotService.InputControllerLastEvents = {
  methodName: "InputControllerLastEvents",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.InputControllerLastEventsRequest,
  responseType: proto_api_v1_robot_pb.InputControllerLastEventsResponse
};

RobotService.InputControllerEventStream = {
  methodName: "InputControllerEventStream",
  service: RobotService,
  requestStream: false,
  responseStream: true,
  requestType: proto_api_v1_robot_pb.InputControllerEventStreamRequest,
  responseType: proto_api_v1_robot_pb.InputControllerEvent
};

RobotService.InputControllerInjectEvent = {
  methodName: "InputControllerInjectEvent",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.InputControllerInjectEventRequest,
  responseType: proto_api_v1_robot_pb.InputControllerInjectEventResponse
};

RobotService.ResourceRunCommand = {
  methodName: "ResourceRunCommand",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ResourceRunCommandRequest,
  responseType: proto_api_v1_robot_pb.ResourceRunCommandResponse
};

RobotService.FrameServiceConfig = {
  methodName: "FrameServiceConfig",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.FrameServiceConfigRequest,
  responseType: proto_api_v1_robot_pb.FrameServiceConfigResponse
};

RobotService.NavigationServiceMode = {
  methodName: "NavigationServiceMode",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceModeRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceModeResponse
};

RobotService.NavigationServiceSetMode = {
  methodName: "NavigationServiceSetMode",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceSetModeRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceSetModeResponse
};

RobotService.NavigationServiceLocation = {
  methodName: "NavigationServiceLocation",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceLocationRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceLocationResponse
};

RobotService.NavigationServiceWaypoints = {
  methodName: "NavigationServiceWaypoints",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceWaypointsRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceWaypointsResponse
};

RobotService.NavigationServiceAddWaypoint = {
  methodName: "NavigationServiceAddWaypoint",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceAddWaypointRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceAddWaypointResponse
};

RobotService.NavigationServiceRemoveWaypoint = {
  methodName: "NavigationServiceRemoveWaypoint",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.NavigationServiceRemoveWaypointRequest,
  responseType: proto_api_v1_robot_pb.NavigationServiceRemoveWaypointResponse
};

RobotService.ObjectManipulationServiceDoGrab = {
  methodName: "ObjectManipulationServiceDoGrab",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.ObjectManipulationServiceDoGrabRequest,
  responseType: proto_api_v1_robot_pb.ObjectManipulationServiceDoGrabResponse
};

RobotService.GPSLocation = {
  methodName: "GPSLocation",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.GPSLocationRequest,
  responseType: proto_api_v1_robot_pb.GPSLocationResponse
};

RobotService.GPSAltitude = {
  methodName: "GPSAltitude",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.GPSAltitudeRequest,
  responseType: proto_api_v1_robot_pb.GPSAltitudeResponse
};

RobotService.GPSSpeed = {
  methodName: "GPSSpeed",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.GPSSpeedRequest,
  responseType: proto_api_v1_robot_pb.GPSSpeedResponse
};

RobotService.GPSAccuracy = {
  methodName: "GPSAccuracy",
  service: RobotService,
  requestStream: false,
  responseStream: false,
  requestType: proto_api_v1_robot_pb.GPSAccuracyRequest,
  responseType: proto_api_v1_robot_pb.GPSAccuracyResponse
};

exports.RobotService = RobotService;

function RobotServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

RobotServiceClient.prototype.status = function status(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.Status, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.statusStream = function statusStream(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(RobotService.StatusStream, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.config = function config(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.Config, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.doAction = function doAction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.DoAction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.baseMoveStraight = function baseMoveStraight(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BaseMoveStraight, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.baseMoveArc = function baseMoveArc(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BaseMoveArc, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.baseSpin = function baseSpin(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BaseSpin, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.baseStop = function baseStop(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BaseStop, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.baseWidthMillis = function baseWidthMillis(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BaseWidthMillis, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardStatus = function boardStatus(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardStatus, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardGPIOSet = function boardGPIOSet(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardGPIOSet, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardGPIOGet = function boardGPIOGet(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardGPIOGet, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardPWMSet = function boardPWMSet(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardPWMSet, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardPWMSetFrequency = function boardPWMSetFrequency(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardPWMSetFrequency, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardAnalogReaderRead = function boardAnalogReaderRead(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardAnalogReaderRead, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardDigitalInterruptConfig = function boardDigitalInterruptConfig(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardDigitalInterruptConfig, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardDigitalInterruptValue = function boardDigitalInterruptValue(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardDigitalInterruptValue, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.boardDigitalInterruptTick = function boardDigitalInterruptTick(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.BoardDigitalInterruptTick, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.sensorReadings = function sensorReadings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.SensorReadings, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.compassHeading = function compassHeading(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.CompassHeading, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.compassStartCalibration = function compassStartCalibration(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.CompassStartCalibration, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.compassStopCalibration = function compassStopCalibration(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.CompassStopCalibration, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.compassMark = function compassMark(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.CompassMark, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.forceMatrixMatrix = function forceMatrixMatrix(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ForceMatrixMatrix, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.forceMatrixSlipDetection = function forceMatrixSlipDetection(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ForceMatrixSlipDetection, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.executeFunction = function executeFunction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ExecuteFunction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.executeSource = function executeSource(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ExecuteSource, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.inputControllerControls = function inputControllerControls(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.InputControllerControls, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.inputControllerLastEvents = function inputControllerLastEvents(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.InputControllerLastEvents, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.inputControllerEventStream = function inputControllerEventStream(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(RobotService.InputControllerEventStream, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.inputControllerInjectEvent = function inputControllerInjectEvent(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.InputControllerInjectEvent, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.resourceRunCommand = function resourceRunCommand(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ResourceRunCommand, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.frameServiceConfig = function frameServiceConfig(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.FrameServiceConfig, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceMode = function navigationServiceMode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceMode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceSetMode = function navigationServiceSetMode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceSetMode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceLocation = function navigationServiceLocation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceLocation, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceWaypoints = function navigationServiceWaypoints(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceWaypoints, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceAddWaypoint = function navigationServiceAddWaypoint(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceAddWaypoint, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.navigationServiceRemoveWaypoint = function navigationServiceRemoveWaypoint(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.NavigationServiceRemoveWaypoint, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.objectManipulationServiceDoGrab = function objectManipulationServiceDoGrab(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.ObjectManipulationServiceDoGrab, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.gPSLocation = function gPSLocation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.GPSLocation, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.gPSAltitude = function gPSAltitude(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.GPSAltitude, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.gPSSpeed = function gPSSpeed(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.GPSSpeed, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RobotServiceClient.prototype.gPSAccuracy = function gPSAccuracy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RobotService.GPSAccuracy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.RobotServiceClient = RobotServiceClient;

