<script setup lang="ts">
import DispNoteIntro from './components/HelloDispNote.vue'
import NoteCreator from './components/NoteCreator.vue'
</script>

<template>
  <header>
    <div class="wrapper">
      <DispNoteIntro msg="DispNote" sub="Create self destructing, encrypted messages" />
    </div>
  </header>

  <main>
    <NoteCreator />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}
header h1{
  font-size: 3em;
}

main{
  margin-top: 2em;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
