
function formatDuration(duration) {
    duration = Math.round(duration / 1000);
    const w = Math.floor(duration / (7 * 24 * 60 * 60));
    duration -= w * (7 * 24 * 60 * 60);
    const d = Math.floor(duration / (24 * 60 * 60));
    duration -= d * (24 * 60 * 60);
    const h = Math.floor(duration / (60 * 60));
    duration -= h * (60 * 60);
    const m = Math.floor(duration / 60);
    duration -= m * 60;
    const s = duration;

    if (w > 0) {
        return `${w}w${d}d${h}h${m}m${s}s`;
    }
    if (d > 0) {
        return `${d}d${h}h${m}m${s}s`;
    }
    if (h > 0) {
        return `${h}h${m}m${s}s`;
    }
    if (m > 0) {
        return `${m}m${s}s`;
    }
    return `${s}s`;
}

function applyCountdown(element) {
    const duration = element.dataset.duration;
    let end = (new Date()).getTime() + duration * 1000;
    const interval = setInterval(function () {
        const timeLeft = end - (new Date()).getTime();
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