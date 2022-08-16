<script setup lang="ts">
import ReadNoteIntro from './components/HelloReadNote.vue'
import NoteViewer from './components/NoteViewer.vue'
</script>

<template>
  <header>
    <h1>📝</h1>
    <div class="wrapper">
      <ReadNoteIntro msg="ReadNote" sub="Someone has a note for you!" />
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
