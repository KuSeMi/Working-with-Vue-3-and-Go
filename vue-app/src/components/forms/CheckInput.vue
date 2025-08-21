<template>
  <div class="form-check">
          <input class="form-check-input" :checked="checked" :required="required" type="checkbox" :value="value" :id="name" :name="name">
          <label class="form-check-label" for="name">
            {{ label }}
          </label>
        </div>
</template>

<script>
export default {
  name: 'CheckInput',
  rops: ["label", "required", "name", "value", "checked"],
}
</script>