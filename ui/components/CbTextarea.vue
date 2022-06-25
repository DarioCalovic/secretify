<template>
  <div class="dark mb-3">{{ text }}</div>
</template>
<style>
.dark {
  border: 0;
  font-family: 'Courier New', Courier, monospace;
  font-size: 1.2rem;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
<script>
export default {
  props: {
    text: String,
  },
}
</script>
