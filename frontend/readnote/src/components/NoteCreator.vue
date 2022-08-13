<script setup lang="ts">

</script>

<script>
export default {
  data() {
    return {
      message: '',
      noteCreateResponse: ''
    }
  },
methods: {
  async submitNote(event) {
    // `this` inside methods points to the current active instance
    //alert(`Hello ${this.name}!`)
    // `event` is the native DOM event

    let getNoteID = async function(){
      let id = ""
      let arr = new Uint8Array(1)

      while (id.length < 22){
        crypto.getRandomValues(arr)
        if ("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_".includes(String.fromCharCode(arr[0]))){
          id += String.fromCharCode(arr[0])
        }
      }
      return id
    }

    if (event) {
      let noteID = await getNoteID()
      let rootURL = this.backendURL
      let backendURL = this.backendURL + '/put/' + noteID

      fetch(backendURL, {
        method: 'PUT',
        body: this.$data.message,
      })
      .then((result) => {
        this.$data.noteCreateResponse = ""
        let error = false
        if (result.status == 200){
          this.$data.noteCreateResponse = "Success: "
        }
        else{
          this.$data.noteCreateResponse = "Error: "
          error = true
        }
        result.text().then((text) =>{
          if (! error){
            this.$data.noteCreateResponse += (text + " Copy this URL: " + rootURL + "/get/" + noteID)
          }
          else{
            this.$data.noteCreateResponse += text
          }
        })
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
  <span class="response">{{ noteCreateResponse }}</span>
</form>

</template>

<style scoped>
.response{
  margin-left: 1em;
}
textarea{
  width: 100%;
  min-height: 350px;
  margin-bottom: 1em;
}
</style>