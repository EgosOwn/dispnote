<script setup lang="ts">
</script>

<script lang="ts">
import aes from 'crypto-js/aes';
import enc from 'crypto-js/enc-utf8';

export default {
  mounted(){
    this.downloadNote()
  },
  data() {
    return {
      decryptedNote: ""
    }
  },
methods: {

  async downloadNote(){
    let noteID = document.location.hash.replace('#', '').substring(0, 31)
    let rootURL = this.backendURL
    let backendURL = this.backendURL + '/get/' + noteID
    fetch(backendURL, {method: 'GET'}).then((result) => {

      let error = false
      if (result.status == 200){
        this.decryptNote(result)
      }
      else{
        error = true
        result.text().then((errMsg) => {
          this.$data.decryptedNote = errMsg
        })
      }
      console.log('Success:', result)
    })
    .catch((error) => {
      console.error('Error:', error)
    })
  },
  async decryptNote(Response){
    let body = await Response.text()
    let key = document.location.hash.replace('#', '').substring(31)
    console.debug(body)
    let decrypted = aes.decrypt(body, key).toString(enc)
    if (decrypted.length > 0){
      this.$data.decryptedNote = decrypted
    }
    else{
      this.$data.decryptedNote = "Error: The note is either corrupted or you have an invalid encryption key.\n\n" +
      "Please report this if you believe this is an error."
    }
    //console.debug(this.$data.decryptedNote)
  }
}
}
</script>
<template>

<textarea readonly v-model="decryptedNote"></textarea>

</template>

<style scoped>
textarea{
  width: 100%;
  min-height: 350px;
  margin-bottom: 1em;
}
</style>