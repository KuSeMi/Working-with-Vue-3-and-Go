<template>
  <div class="container" v-if="ready">
    <div class="row">
      <div class="col-md-2">
        <img :src="`${imgPath}/covers/${book.slug}.jpg`" alt="cover" class="img-fluid img-thumbnail">
      </div>
      <div class="col-md-10">
        <h3 class="mt-3">
          {{ book.title }}<hr>
        </h3>
        <p>
          <strong>Author:</strong>{{ book.author?.author_name }}<br>
          <strong>Published:</strong>{{ book.publication_year }}
        </p>
        <p>
          {{ book.description  }}
        </p>
      </div>
    </div>
  </div>

  <p v-else>Loading...</p>
</template>

<script>
export default {
  name: 'AppBook',

  data() {
    return {
      book: {},
      ready: false,
      imgPath: process.env.VUE_APP_IMAGE_URL,
    }
  },
  mounted() {
    fetch(process.env.VUE_APP_API_URL + '/books/' + this.$route.params.bookName)
      .then(response => response.json())
      .then((data) => {
        if (data.error) {
          this.$emit('error', data.message);
        } else {
          this.book = data.data;
          this.ready = true;
        }
      })
  },
}
</script>