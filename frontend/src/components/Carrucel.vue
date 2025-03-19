<script setup>
import { ref, computed } from "vue";
import Slide from "./Slide.vue";

const items = [
  { id: 1, title: "Slide 1", description: "Descripción del Slide 1", image: "https://via.placeholder.com/150" },
  { id: 2, title: "Slide 2", description: "Descripción del Slide 2", image: "https://via.placeholder.com/150" },
  { id: 3, title: "Slide 3", description: "Descripción del Slide 3", image: "https://via.placeholder.com/150" },
  { id: 4, title: "Slide r", description: "Descripción del Slide 3", image: "https://via.placeholder.com/150" },
];

const currentIndex = ref(0);

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % items.length;
};

const prevSlide = () => {
  currentIndex.value = (currentIndex.value - 1 + items.length) % items.length;
};

const slideStyle = computed(() => ({
  transform: `translateX(-${currentIndex.value * 100}%)`,
}));
</script>

<template>
  <div class="relative w-full overflow-hidden">
    <!-- Carrusel -->
    <div class="flex transition-transform duration-500 ease-in-out" :style="slideStyle">
      <Slide
        v-for="item in items"
        :key="item.id"
        :title="item.title"
        :description="item.description"
        :image="item.image"
      />
    </div>

    <!-- Botones -->
    <button
      @click="prevSlide"
      class="absolute top-1/2 left-2 transform -translate-y-1/2 bg-gray-700 text-white p-2 rounded-full"
    >
      ◀
    </button>

    <button
      @click="nextSlide"
      class="absolute top-1/2 right-2 transform -translate-y-1/2 bg-gray-700 text-white p-2 rounded-full"
    >
      ▶
    </button>
  </div>
</template>
