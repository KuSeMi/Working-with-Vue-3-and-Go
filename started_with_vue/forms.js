const TextInput = {
      props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: String,
        min: String,
        max: String,
        value: String,
      },
      template: `
        <div class="mb-3">
          <label :for="name" class="form-label">{{ label }}</label>
          <input
            :type="type"
            :name="name"
            :placeholder="placeholder"
            :required="required"
            :min="min"
            :max="max"
            :value="value"
            :autocomplete="name + '-new'"
            class="form-control"
          >
        </div>
      `
    }

    const SelectInput = {
      props: ["items", "name", "required", "label"],
      template: `
        <div class="mb-3">
          <label :for="name" class="form-label">{{ label }}</label>
          <select :id="name" :name="name" class="form-select" :required="required">
            <option v-for="option in items" :value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>
      `
    }

    const CheckInput = {
      props: ["label", "required", "name", "value", "checked"],
      template: `
        <div class="form-check">
          <input class="form-check-input" :checked="checked" :required="required" type="checkbox" :value="value" :id="name" :name="name">
          <label class="form-check-label" for="name">
            {{ label }}
          </label>
        </div>
      `
    }