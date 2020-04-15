'use strict';

function formatDuration(duration) {
    duration = Math.round(duration / 1000);
    var w = Math.floor(duration / (7 * 24 * 60 * 60));
    duration -= w * (7 * 24 * 60 * 60);
    var d = Math.floor(duration / (24 * 60 * 60));
    duration -= d * (24 * 60 * 60);
    var h = Math.floor(duration / (60 * 60));
    duration -= h * (60 * 60);
    var m = Math.floor(duration / 60);
    duration -= m * 60;
    var s = duration;

    if (w > 0) {
        return w + 'w' + d + 'd' + h + 'h' + m + 'm' + s + 's';
    }
    if (d > 0) {
        return d + 'd' + h + 'h' + m + 'm' + s + 's';
    }
    if (h > 0) {
        return h + 'h' + m + 'm' + s + 's';
    }
    if (m > 0) {
        return m + 'm' + s + 's';
    }
    return s + 's';
}

function applyCountdown(element) {
    var duration = element.dataset.duration;
    var end = new Date().getTime() + duration * 1000;
    var interval = setInterval(function () {
        var timeLeft = end - new Date().getTime();
        if (timeLeft < 0) {
            element.innerText = 'Done';
            document.getElementsByClassName('Planet-card__button--cancel')[0].disabled = true;
            clearInterval(interval);
            return;
        }
        element.innerText = formatDuration(timeLeft);
    }, 1000);
}

function initCountdowns() {
    document.querySelectorAll('[data-duration]').forEach(function (element) {
        applyCountdown(element);
    });
}

document.addEventListener('DOMContentLoaded', initCountdowns, false);
