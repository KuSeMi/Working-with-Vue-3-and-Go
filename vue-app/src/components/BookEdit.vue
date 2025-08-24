<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Book Edit</h1>
        <hr>
      </div>
    </div>
  </div>
</template>

<script>
import Security from './security';

export default {
  name: 'BookEdit',

  beforeMount() {
    Security.requireToken();
  }
}
</script>