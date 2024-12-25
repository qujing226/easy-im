/** 数组转化为字符串*/
export function arrayToString(arr, field, isTrans) {
    isTrans = typeof isTrans !== 'undefined' ? isTrans : false;
    var idr = [];
    for (var i = 0; i < arr.length; i++) {
        idr.push(arr[i][field]);
    }
    if (isTrans === true) {
        return idr.join(",");
    } else {
        return idr;
    }
}
/** 删除数组中某一个值*/
export function delArrValue(arr, field, value) {
    for (var i = 0; i < arr.length; i++) {
        if (arr[i][field] == value) {
            arr.splice(i, 1);
        }
    }
    return arr;
}
export function timeFormat(t) {
    var time = parseInt(t / 1000);
    return isToday(t) ? date("H:i", time) : date("Y/m/d", time);
}
export function date(format, timestamp) {
    var a, jsdate = ((timestamp) ? new Date(timestamp * 1000) : new Date());
    var pad = function(n, c) {
        if ((n = n + "").length < c) {
            return new Array(++c - n.length).join("0") + n;
        } else {
            return n;
        }
    };
    var txt_weekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    var txt_ordin = { 1: "st", 2: "nd", 3: "rd", 21: "st", 22: "nd", 23: "rd", 31: "st" };
    var txt_months = ["", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    var f = {
        d: function() { return pad(f.j(), 2) },
        D: function() { return f.l().substr(0, 3) },
        j: function() { return jsdate.getDate() },
        l: function() { return txt_weekdays[f.w()] },
        N: function() { return f.w() + 1 },
        S: function() { return txt_ordin[f.j()] ? txt_ordin[f.j()] : 'th' },
        w: function() { return jsdate.getDay() },
        z: function() { return (jsdate - new Date(jsdate.getFullYear() + "/1/1")) / 864e5 >> 0 },
        W: function() {
            var a = f.z(),
                b = 364 + f.L() - a;
            var nd2, nd = (new Date(jsdate.getFullYear() + "/1/1").getDay() || 7) - 1;
            if (b <= 2 && ((jsdate.getDay() || 7) - 1) <= 2 - b) {
                return 1;
            } else {
                if (a <= 2 && nd >= 4 && a >= (6 - nd)) {
                    nd2 = new Date(jsdate.getFullYear() - 1 + "/12/31");
                    return date("W", Math.round(nd2.getTime() / 1000));
                } else {
                    return (1 + (nd <= 3 ? ((a + nd) / 7) : (a - (7 - nd)) / 7) >> 0);
                }
            }
        },
        F: function() { return txt_months[f.n()] },
        m: function() { return pad(f.n(), 2) },
        M: function() { return f.F().substr(0, 3) },
        n: function() { return jsdate.getMonth() + 1 },
        t: function() {
            var n;
            if ((n = jsdate.getMonth() + 1) === 2) {
                return 28 + f.L();
            } else {
                if (n & 1 && n < 8 || !(n & 1) && n > 7) {
                    return 31;
                } else {
                    return 30;
                }
            }
        },
        L: function() { var y = f.Y(); return (!(y & 3) && (y % 1e2 || !(y % 4e2))) ? 1 : 0 },
        Y: function() { return jsdate.getFullYear() },
        y: function() { return (jsdate.getFullYear() + "").slice(2) },
        a: function() { return jsdate.getHours() > 11 ? "pm" : "am" },
        A: function() { return f.a().toUpperCase() },
        B: function() {
            var off = (jsdate.getTimezoneOffset() + 60) * 60;
            var theSeconds = (jsdate.getHours() * 3600) + (jsdate.getMinutes() * 60) + jsdate.getSeconds() + off;
            var beat = Math.floor(theSeconds / 86.4);
            if (beat > 1000) beat -= 1000;
            if (beat < 0) beat += 1000;
            if ((String(beat)).length === 1) beat = "00" + beat;
            if ((String(beat)).length === 2) beat = "0" + beat;
            return beat;
        },
        g: function() { return jsdate.getHours() % 12 || 12 },
        G: function() { return jsdate.getHours() },
        h: function() { return pad(f.g(), 2) },
        H: function() { return pad(jsdate.getHours(), 2) },
        i: function() { return pad(jsdate.getMinutes(), 2) },
        s: function() { return pad(jsdate.getSeconds(), 2) },
        O: function() {
            var t = pad(Math.abs(jsdate.getTimezoneOffset() / 60 * 100), 4);
            if (jsdate.getTimezoneOffset() > 0) t = "-" + t;
            else t = "+" + t;
            return t;
        },
        P: function() { var O = f.O(); return (O.substr(0, 3) + ":" + O.substr(3, 2)) },
        c: function() { return f.Y() + "-" + f.m() + "-" + f.d() + "T" + f.h() + ":" + f.i() + ":" + f.s() + f.P() },
        U: function() { return Math.round(jsdate.getTime() / 1000) }
    };
    let ret = '';
    return format.replace(/[\\]?([a-zA-Z])/g, function(t, s) {
        if (t !== s) {
            ret = s;
        } else if (f[s]) {
            ret = f[s]();
        } else {
            ret = s;
        }
        return ret;
    });
}
// 判断是否为今天
export function isToday(time) {
    // 获取当前日期
    const today = new Date().setHours(0, 0, 0, 0); // 将时分秒毫秒设为0
    // 判断是否是今天
    if (time >= today && time < today + 24 * 3600 * 1000) {
        return true;
    } else {
        return false;
    }
}

export function generateRandId(){
	var d = new Date().getTime();
	  var uuid = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
	    var r = (d + Math.random() * 16) % 16 | 0;
	    d = Math.floor(d / 16);
	    return (c == "x" ? r : r & 0x3 | 0x8).toString(16);
	  });
	  return uuid;
}
