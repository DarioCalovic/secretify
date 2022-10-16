<template>
  <div class="dark mb-3" v-if="first">
    <span class="first">{{ first }}</span
    ><span>{{ rest }}</span>
  </div>
  <div class="dark mb-3" v-else>{{ text }}</div>
</template>
<style>
.dark {
  border: 0;
  font-family: 'Courier New', Courier, monospace;
  font-size: 1.2rem;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: pre-line;
  min-height: 3.1rem;
}
</style>
<script>
export default {
  props: {
    text: String,
    first: String,
    rest: String,
  },
}
</script>
