# Vue.js Framework Learning Examples / Приклади для вивчення фреймворка Vue.js

## English / Англійська

### 🎯 Overview

This directory contains Vue.js learning examples that demonstrate the most important concepts of the Vue.js framework. These examples are built using Vue 3 with the Options API and showcase fundamental Vue concepts through practical implementations.

### 📁 Files Structure

- **`index.html`** - Main example demonstrating core Vue concepts
- **`fetch.html`** - Example showing data fetching and component communication
- **`registration.html`** - Registration form using custom components
- **`forms.js`** - Reusable form components library
- **`registration-form.js`** - Complex form component with conditional rendering

### 🔑 Key Vue.js Concepts Demonstrated

#### 1. **Reactive Data Binding**
Vue's reactive data system automatically updates the DOM when data changes.

```javascript
// Example from index.html
data() {
  return {
    message: "Hello Vue!",
    counter: 0
  }
}
```

**Template usage:**
```html
<h1>{{ message }}</h1>
<h2>{{ counter }}</h2>
```

#### 2. **Directives**
Special HTML attributes that provide dynamic behavior to elements.

- **`v-model`** - Two-way data binding for form inputs
- **`v-on` / `@`** - Event listeners
- **`v-for`** - List rendering
- **`v-if`** - Conditional rendering
- **`v-bind` / `:`** - Attribute binding

```html
<!-- Two-way binding -->
<input v-model="message" type="text" class="form-control">

<!-- Event handling -->
<button @click="increment">Increment</button>

<!-- List rendering -->
<book-item v-for="item in books.results" :key="item.id" :book="item"></book-item>

<!-- Conditional rendering -->
<div v-if="addressSameChecked === false">
  <!-- Conditional content -->
</div>
```

#### 3. **Component System**
Reusable, self-contained pieces of code that encapsulate HTML, CSS, and JavaScript logic.

```javascript
// Component definition
const TextInput = {
  props: {
    name: String,
    type: String,
    label: String,
    required: String
  },
  template: `
    <div class="mb-3">
      <label :for="name" class="form-label">{{ label }}</label>
      <input :type="type" :name="name" :required="required" class="form-control">
    </div>
  `
}

// Component registration
components: {
  TextInput,
  'text-input': TextInput
}
```

#### 4. **Props (Properties)**
Mechanism for passing data from parent to child components.

```javascript
// Parent component passes data
<text-input label="First Name" name="first-name" type="text" required="true"></text-input>

// Child component receives data
props: {
  label: String,
  name: String,
  type: String,
  required: String
}
```

#### 5. **Event Handling & Communication**
Components can emit events to communicate with parent components.

```javascript
// Child emits event
<a href="#!" @click="$emit('removebook', book.id)">Delete</a>

// Parent listens to event
<book-item @removeBook="removeBook"></book-item>
```

#### 6. **Lifecycle Hooks**
Methods that are called at specific stages of a component's lifecycle.

```javascript
// Component lifecycle hooks
mounted() {
  // Called after component is mounted to DOM
  console.log("Component mounted");
  
  // Fetch data from API
  fetch("https://gutendex.com/books/")
    .then(response => response.json())
    .then(data => this.books = data);
},

created() {
  // Called when component is created
},

updated() {
  // Called after component data changes
}
```

#### 7. **Methods**
Functions that can be called from templates or other methods.

```javascript
methods: {
  increment() {
    this.counter++;
  },
  
  removeBook(id) {
    this.books.results = this.books.results.filter(book => book.id !== id);
  },
  
  addressSame() {
    this.addressSameChecked = !this.addressSameChecked;
  }
}
```

#### 8. **Computed Properties & Watchers**
Reactive computed values and data change watchers (demonstrated implicitly in the examples).

### 🚀 Running the Examples

1. Open any HTML file in a web browser
2. The examples use Vue 3 from CDN, so no build process is required
3. Bootstrap CSS is included for styling

### 💡 Learning Path Recommendations

1. Start with `index.html` to understand basic concepts
2. Explore `fetch.html` for API integration and component communication  
3. Study `registration.html` and related JS files for complex forms and component composition

---

## Українська / Ukrainian

### 🎯 Огляд

Ця директорія містить навчальні приклади Vue.js, які демонструють найважливіші концепції фреймворка Vue.js. Ці приклади створені з використанням Vue 3 з Options API і показують фундаментальні концепції Vue через практичні реалізації.

### 📁 Структура файлів

- **`index.html`** - Головний приклад, що демонструє основні концепції Vue
- **`fetch.html`** - Приклад показу отримання даних та комунікації між компонентами
- **`registration.html`** - Форма реєстрації з використанням кастомних компонентів
- **`forms.js`** - Бібліотека багаторазових компонентів форм
- **`registration-form.js`** - Складний компонент форми з умовним рендерингом

### 🔑 Ключові концепції Vue.js, що демонструються

#### 1. **Реактивна прив'язка даних**
Реактивна система даних Vue автоматично оновлює DOM при зміні даних.

```javascript
// Приклад з index.html
data() {
  return {
    message: "Hello Vue!",
    counter: 0
  }
}
```

**Використання в шаблоні:**
```html
<h1>{{ message }}</h1>
<h2>{{ counter }}</h2>
```

#### 2. **Директиви**
Спеціальні HTML атрибути, що надають динамічну поведінку елементам.

- **`v-model`** - Двостороння прив'язка даних для елементів форм
- **`v-on` / `@`** - Слухачі подій
- **`v-for`** - Рендеринг списків
- **`v-if`** - Умовний рендеринг
- **`v-bind` / `:`** - Прив'язка атрибутів

```html
<!-- Двостороння прив'язка -->
<input v-model="message" type="text" class="form-control">

<!-- Обробка подій -->
<button @click="increment">Збільшити</button>

<!-- Рендеринг списку -->
<book-item v-for="item in books.results" :key="item.id" :book="item"></book-item>

<!-- Умовний рендеринг -->
<div v-if="addressSameChecked === false">
  <!-- Умовний контент -->
</div>
```

#### 3. **Система компонентів**
Багаторазові, самодостатні частини коду, що інкапсулюють HTML, CSS та JavaScript логіку.

```javascript
// Визначення компонента
const TextInput = {
  props: {
    name: String,
    type: String,
    label: String,
    required: String
  },
  template: `
    <div class="mb-3">
      <label :for="name" class="form-label">{{ label }}</label>
      <input :type="type" :name="name" :required="required" class="form-control">
    </div>
  `
}

// Реєстрація компонента
components: {
  TextInput,
  'text-input': TextInput
}
```

#### 4. **Пропси (Властивості)**
Механізм передачі даних від батьківського до дочірнього компонента.

```javascript
// Батьківський компонент передає дані
<text-input label="Ім'я" name="first-name" type="text" required="true"></text-input>

// Дочірній компонент отримує дані
props: {
  label: String,
  name: String,
  type: String,
  required: String
}
```

#### 5. **Обробка подій і комунікація**
Компоненти можуть випромінювати події для комунікації з батьківськими компонентами.

```javascript
// Дочірній компонент випромінює подію
<a href="#!" @click="$emit('removebook', book.id)">Видалити</a>

// Батьківський компонент слухає подію
<book-item @removeBook="removeBook"></book-item>
```

#### 6. **Хуки життєвого циклу**
Методи, що викликаються на специфічних етапах життєвого циклу компонента.

```javascript
// Хуки життєвого циклу компонента
mounted() {
  // Викликається після монтування компонента в DOM
  console.log("Компонент змонтовано");
  
  // Отримання даних з API
  fetch("https://gutendex.com/books/")
    .then(response => response.json())
    .then(data => this.books = data);
},

created() {
  // Викликається при створенні компонента
},

updated() {
  // Викликається після зміни даних компонента
}
```

#### 7. **Методи**
Функції, що можуть викликатися з шаблонів або інших методів.

```javascript
methods: {
  increment() {
    this.counter++;
  },
  
  removeBook(id) {
    this.books.results = this.books.results.filter(book => book.id !== id);
  },
  
  addressSame() {
    this.addressSameChecked = !this.addressSameChecked;
  }
}
```

#### 8. **Обчислені властивості і спостерігачі**
Реактивні обчислені значення та спостерігачі змін даних (демонструються неявно в прикладах).

### 🚀 Запуск прикладів

1. Відкрийте будь-який HTML файл у веб-браузері
2. Приклади використовують Vue 3 з CDN, тому процес збірки не потрібен
3. Bootstrap CSS включено для стилізації

### 💡 Рекомендації щодо навчання

1. Почніть з `index.html` для розуміння базових концепцій
2. Вивчіть `fetch.html` для інтеграції з API та комунікації між компонентами
3. Дослідіть `registration.html` та пов'язані JS файли для складних форм та композиції компонентів

---

## 📚 Additional Resources / Додаткові ресурси

- [Vue.js Official Documentation](https://vuejs.org/)
- [Vue.js Офіційна документація (українською)](https://ua.vuejs.org/)
- [Vue 3 Migration Guide](https://v3-migration.vuejs.org/)
- [Vue.js Examples and Tutorials](https://vuejsexamples.com/)

