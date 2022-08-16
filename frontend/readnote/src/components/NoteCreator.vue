<script setup lang="ts">

</script>

<script>
export default {
  data() {
    return {
      message: '',
      noteCreateResponse: '',
      noteCreatedURL: '',
      formShow: true,
      key: new ArrayBuffer(),
      noteID: ''
    }
  },
methods: {
  async createNote(){
      return window.crypto.subtle.generateKey(
          {
              name: "AES-CTR",
              length: 256, //can be  128, 192, or 256
          },
          true, //whether the key is extractable (i.e. can be used in exportKey)
          ["encrypt", "decrypt"] //can "encrypt", "decrypt", "wrapKey", or "unwrapKey"
      )
      .then((key)=>{
          //returns a key object
          window.crypto.subtle.exportKey("raw", key).then((exportedKey)=>{
            this.$data.noteID += "#" + btoa(exportedKey)
          })
          let encoder = new TextEncoder()

           window.crypto.subtle.encrypt(
              {
                  name: "AES-CTR",
                  //Don't re-use counters!
                  //Always use a new counter every time your encrypt!
                  counter: new Uint8Array(16),
                  length: 128, //can be 1-128
              },
              key, //from generateKey or importKey above
              encoder.encode(this.$data.message) //ArrayBuffer of data you want to encrypt
          )
          .then((encrypted)=>{
              //returns an ArrayBuffer containing the encrypted data
              this.$data.message = encrypted
              console.log(this.$data.message);
          })
          .catch(function(err){
              console.error(err);
          });
      })
  },
  async submitNote(event) {
    // Prevent the browser from doing a regular form submit
      event.preventDefault()
    // `this` inside methods points to the current active instance
    //alert(`Hello ${this.name}!`)
    // `event` is the native DOM event

    let getNoteID = async function(){
      let id = ""
      let arr = new Uint8Array(8)
      crypto.getRandomValues(arr)
      return btoa(arr).substring(0, 32)
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
            this.$data.noteCreatedURL = rootURL + "/get/" + this.$data.noteID
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