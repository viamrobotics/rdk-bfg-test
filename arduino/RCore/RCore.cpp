
#include "RCore.h"

extern HardwareSerial* debugSerial;

Motor::Motor(int in1, int in2, int pwm)
    : _in1(in1), _in2(in2), _pwm(pwm), _encoderTicks(0), _encoderTicksStop(0) {
    pinMode(_in1, OUTPUT);
    pinMode(_in2, OUTPUT);
    pinMode(_pwm, OUTPUT);
    _moving = false;
}

void Motor::stop() {
    _moving = false;
    _encoderTicksStop = 0;
    digitalWrite(_in1, LOW);
    digitalWrite(_in2, LOW);
}

void Motor::forward(int val) {
    _moving = true;
    analogWrite(_pwm, val);
    digitalWrite(_in1, HIGH);
    digitalWrite(_in2, LOW);
}

void Motor::backward(int val) {
    _moving = true;
    analogWrite(_pwm, val);
    digitalWrite(_in1, LOW);
    digitalWrite(_in2, HIGH);
}

void Motor::doCommand(const char* buf) {
    Command c = Command::parse(buf);
    if (c.ticks == 0) {
        _encoderTicksStop = 0;
    } else {
        _encoderTicksStop = c.ticks + _encoderTicks;
    }

    switch (c.direction) {
        case 'f':
            forward(c.speed);
            break;
        case 'b':
            backward(c.speed);
            break;
        case 's':
            stop();
            break;
        default:
            debugSerial->println("unknown command");
            debugSerial->println(buf[0], DEC);
    }
}

bool Motor::checkEncoder() {
    if (_encoderTicksStop <= 0) {
        return false;
    }

    if (_encoderTicks > _encoderTicksStop) {
        stop();
        return true;
    }

    if (_encoderTicks + 60 > _encoderTicksStop) {
        // analogWrite(_pwm, 40); // slow down
        return false;
    }

    return false;
}

// -----

Command Command::parse(const char* buf) {
    Command c;

    if (!buf[0]) {
        return c;
    }

    c.direction = buf[0];
    buf++;

    if (!buf[0]) {
        return c;
    }

    c.speed = atoi(buf);
    if (c.speed <= 0 || c.speed > 255) {
        // bad data, do nothing
        c.direction = 's';
        c.speed = 0;
        return c;
    }

    // move pase the number to see if we have more data

    while (isdigit(buf[0])) {
        buf++;
    }

    if (buf[0] != ',') {
        return c;
    }
    buf++;  // move past the comma

    c.ticks = atoi(buf);

    return c;
}

void _testParseCommand(const char* buf, Command correct) {
    Command c = Command::parse(buf);
    if (c.direction == correct.direction && c.speed == correct.speed &&
        c.ticks == correct.ticks) {
        return;
    }

    Serial.println(buf);
    Serial.println("BROKE");
    exit(-1);
}

void testParseCommand() {
    _testParseCommand("s", Command('s', 255, 0));

    _testParseCommand("f", Command('f', 255, 0));
    _testParseCommand("f9", Command('f', 9, 0));
    _testParseCommand("f91", Command('f', 91, 0));
    _testParseCommand("f191", Command('f', 191, 0));
    _testParseCommand("f1000", Command('s', 0, 0));

    _testParseCommand("b91", Command('b', 91, 0));

    _testParseCommand("f100,100", Command('f', 100, 100));
}

// ------

// return true if got a new line
bool Buffer::readTillNewLine() {
    while (_port->available()) {
        int x = _port->read();
        if (x == '\n') {
            continue;
        }

        if (x == '\r') {
            _buf[_pos] = 0;
            return true;
        }

        if (_pos > 200) {
            Serial.println("bad bad");
            return false;
        }

        _buf[_pos++] = x;
    }

    return false;
}

const char* Buffer::getLineAndReset() {
    _buf[_pos] = 0;
    _pos = 0;
    return _buf;
}
