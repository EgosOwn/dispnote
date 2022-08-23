<script setup lang="ts">
import DispNoteIntro from './components/HelloDispNote.vue'
import NoteViewer from './components/NoteViewer.vue'
</script>

<template>
  <header>
    <div class="wrapper">
      <DispNoteIntro msg="DispNote" sub="This message will disappear once the page is closed." />
    </div>
  </header>

  <main>
    <NoteViewer />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}
header h1{
  font-size: 3em;
}

textarea{
  font-size: 1.3em;
}

@media (max-width: 1024px) {
  header{
    margin-bottom: 2em;
  }
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }
  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
