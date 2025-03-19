<!-- Slide.vue -->
<script setup>
import Card from "./Card.vue"; // Importamos el componente Card

defineProps({
  title: String,
  description: String,
  image: String,
});
</script>

<template>
  <div class="w-full flex justify-center items-center">
    <Card />
  </div>
</template>
