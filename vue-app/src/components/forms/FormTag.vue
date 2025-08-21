<template>
  <form 
    @submit.prevent="submit"
    autocomplete="off" 
    :method="method" 
    :action="action" 
    class="needs-validation" 
    novalidate>
    <slot></slot>
  </form>
</template>

<script>
export default {
  name: 'FormTag',
  props: [
    'method',
    'action',
    'event',
    'name'
  ],
  methods: {
    // Bootstrap validation
    submit(event) {
      const myForm = event.target;
      if (myForm.checkValidity()) {
        this.$emit(this.event);
      }
      myForm.classList.add('was-validated');
    }
  },
}
  
</script>