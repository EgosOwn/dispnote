<script setup lang="ts">
</script>

<script lang="ts">
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
    let noteID = document.location.hash.replace('#', '')
    let rootURL = this.backendURL
    let backendURL = this.backendURL + '/get/' + noteID
    fetch(backendURL, {method: 'GET'}).then((result) => {

      let error = false
      if (result.status == 200){
        this.decryptNote(result.body)
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
  async decryptNote(body: ReadableStream){
    body.getReader().read().then((noteData) =>
    
      this.$data.decryptedNote = noteData.value
    )
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