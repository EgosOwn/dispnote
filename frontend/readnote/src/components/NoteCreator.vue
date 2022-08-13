<script setup lang="ts">

</script>

<script>
export default {
  data() {
    return {
      message: ''
    }
  },
methods: {
  submitNote(event) {
    // `this` inside methods points to the current active instance
    //alert(`Hello ${this.name}!`)
    // `event` is the native DOM event
    if (event) {

      fetch('/test', {
        method: 'POST',
        body: this.$data.message,
      })
      .then((response) => response.json())
      .then((result) => {
        console.log('Success:', result);
      })
      .catch((error) => {
        console.error('Error:', error);
      });

      // Prevent the browser from doing a regular form submit
      event.preventDefault()
    }

    return false
  }
}

}
</script>
<template>

<form method="POST" @submit="submitNote">
  <textarea v-model="message" placeholder="the note text" required></textarea>
  <input type="submit" value="Submit Note">
</form>

</template>

<style scoped>
textarea{
  width: 100%;
  min-height: 350px;
  margin-bottom: 1em;
}
</style>