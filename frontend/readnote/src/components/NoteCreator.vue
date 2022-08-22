<script setup lang="ts">

</script>

<script>
import aes from 'crypto-js/aes'
import enc from 'crypto-js/enc-utf8'

export default {
  data() {
    return {
      message: '',
      noteCreateResponse: '',
      noteCreatedURL: '',
      formShow: true,
      key: '',
      noteID: ''
    }
  },
methods: {
  async createNote(){
    let arr = new Int32Array(4); crypto.getRandomValues(arr)

    let key = btoa(arr).replaceAll('=', '')

    this.$data.key = key

    //console.debug(this.$data.message)
    this.$data.message = aes.encrypt(this.$data.message, key)

    //let decrypted = aes.decrypt(this.$data.message, key)
    //console.debug(enc.parse(decrypted))
  },
  async submitNote(event) {
    // Prevent the browser from doing a regular form submit
      event.preventDefault()
    // `this` inside methods points to the current active instance
    //alert(`Hello ${this.name}!`)
    // `event` is the native DOM event

    let getNoteID = async function(){
      let id = ""
      let arr = new Int32Array(2)
      crypto.getRandomValues(arr)
      return btoa(arr).replaceAll('=', '').padStart(31, '0')
    }

    if (event) {
      this.$data.noteID = await getNoteID()
      // Encrypt the note
      await this.createNote()

      let rootURL = this.backendURL
      let backendURL = this.backendURL + '/put/' + this.$data.noteID

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
            this.$data.noteCreateResponse += (text + " Copy this URL: ")
            this.$data.noteCreatedURL = document.location.protocol + '//' + document.location.host + "/read.html#" + this.$data.noteID + this.$data.key
          }
          else{
            this.$data.noteCreateResponse += text
          }
        })
        console.log('Success:', result)
        this.$data.formShow = false
      })
      .catch((error) => {
        console.error('Error:', error)
      })


    }

    return false
  }
}

}
</script>
<template>

<form method="POST" @submit="submitNote" v-if="formShow">
  <textarea v-model="message" placeholder="the note text" required></textarea>
  <input type="submit" value="Submit Note">
</form>
<div v-if="! formShow">
<div>{{ noteCreateResponse }}</div>
<input type="text" v-model="noteCreatedURL" v-if="noteCreatedURL" readonly>
</div>

</template>

<style scoped>
textarea{
  width: 100%;
  min-height: 350px;
  margin-bottom: 1em;
}
</style>