 const RegistrationForm = {
  data() {
    return {
      addressSameChecked: true
    }
  },
  props: ["items"],
  template: `
    <h3>Registration</h3>
    <hr>
    <form autocomplete="off" class="needs-validation" novalidate>
      <text-input label="First Name" name="first-name" type="text" required="true"></text-input>
      <text-input label="Last Name" name="last-name" type="text" required="true"></text-input>
      <text-input label="Email" name="email" type="email" required="true"></text-input>
      <text-input label="Password" name="password" type="password" required="true" ></text-input>
      <select-input label="Favorite Colour" name="color" :items="items"></select-input>
      
      <text-input label="Address" name="address" type="text" required="true"></text-input>
      <text-input label="City/Town" name="city" type="text" required="true"></text-input>
      <text-input label="State/Province" name="province" type="text" required="true"></text-input>
      <text-input label="Zip/Postal" name="zip" type="text" required="true"></text-input>
      
      <check-input v-on:click="addressSame" label="Mailing Address Same" checked="true" v-model="addressSameChecked"></check-input>

      <div v-if="addressSameChecked === false">
        <div class="mb-3">
          <text-input label=" Mailing Address" name="address2" type="text"></text-input>
          <text-input label="Mailing City/Town" name="city2" type="text"></text-input>
          <text-input label="Mailing State/Province" name="province2" type="text"></text-input>
          <text-input label="Mailing Zip/Postal" name="zip2" type="text"></text-input>
        </div>
      </div>

      <check-input label="I agree to terms and conditions" required="true"></check-input>
      <hr>
      <input type="submit" value="Register" class="btn btn-primary">
    </form>
  `,
  methods: {
    addressSame() {
      console.log("address same fired");
      if (this.addressSameChecked === true) {
        console.log("was checked on click")
        this.addressSameChecked = false;
      } else {
        console.log("was not checked on click")
        this.addressSameChecked = true;
      }
    }
  },
  components: {
    'text-input': TextInput,
    'select-input': SelectInput,
    'check-input': CheckInput,
  },
  mounted() {
    // Example starter JavaScript for disabling form submissions if there are invalid fields
    (function () {
      'use strict'

      // Fetch all the forms we want to apply custom Bootstrap validation styles to
      var forms = document.querySelectorAll('.needs-validation')

      // Loop over them and prevent submission
      Array.prototype.slice.call(forms)
        .forEach(function (form) {
          form.addEventListener('submit', function (event) {
            if (!form.checkValidity()) {
              event.preventDefault()
              event.stopPropagation()
            }

            form.classList.add('was-validated')
          }, false)
        })
    })()
  }
 }