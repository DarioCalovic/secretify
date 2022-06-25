<template>
  <div><div id="progressbar1"></div></div>
</template>
<style>
.progressbar .inner {
  height: 15px;
  animation: progressbar-countdown;
  /* Placeholder, this will be updated using javascript */
  animation-duration: 40s;
  /* We stop in the end */
  animation-iteration-count: 1;
  /* Stay on pause when the animation is finished finished */
  animation-fill-mode: forwards;
  /* We start paused, we start the animation using javascript */
  animation-play-state: paused;
  /* We want a linear animation, ease-out is standard */
  animation-timing-function: linear;
}
@keyframes progressbar-countdown {
  0% {
    width: 100%;
    background: #48c78e;
  }
  100% {
    width: 0%;
    background: #f14668;
  }
}
</style>
<script>
/*
 *  Creates a progressbar.
 *  @param id the id of the div we want to transform in a progressbar
 *  @param duration the duration of the timer example: '10s'
 *  @param callback, optional function which is called when the progressbar reaches 0.
 */
function createProgressbar(id, duration, callback) {
  // We select the div that we want to turn into a progressbar
  const progressbar = document.getElementById(id)
  progressbar.className = 'progressbar'

  // We create the div that changes width to show progress
  const progressbarinner = document.createElement('div')
  progressbarinner.className = 'inner'

  // Now we set the animation parameters
  progressbarinner.style.animationDuration = duration

  // Eventually couple a callback
  if (typeof callback === 'function') {
    progressbarinner.addEventListener('animationend', callback)
  }

  // Append the progressbar to the main progressbardiv
  progressbar.appendChild(progressbarinner)

  // When everything is set up we start the animation
  progressbarinner.style.animationPlayState = 'running'
}

export default {
  name: 'ProgressBar',
  mounted() {
    createProgressbar('progressbar1', '30s', function () {
      window.location.reload(true)
    })
  },
}
</script>
