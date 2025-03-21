<template>
    <div>
        <!-- Carrusel -->
        <div class="relative mb-6">
            <div class="relative w-full max-w-3xl mx-auto overflow-hidden z-10">
                <div class="flex transition-transform duration-500 ease-in-out" :style="carouselStore.slideStyle">
                    <div v-for="item in carouselStore.items.filter(item => item.recommended)" :key="item.id" class="w-[calc(100%/3)] flex-shrink-0 px-2">
                        <Card :item="item" />
                    </div>
                </div>
            </div>
            <div class="absolute w-full max-w-3xl mx-auto z-10">
                <button @click="carouselStore.prevSlide" :disabled="carouselStore.currentIndex === 0"
                    class="absolute -top-64 left-52 transform -translate-y-1/2 text-white p-2 rounded-full">
                    <img class="max-h-7 -rotate-90" src="../assets/flecha-hacia-arriba (1).png" alt="">
                </button>

                <button @click="carouselStore.nextSlide" :disabled="carouselStore.items.length === 0"
                    class="absolute -top-64 -right-96 transform -translate-y-1/2 text-white p-2 rounded-full">
                    <img class="max-h-7 rotate-90" src="../assets/flecha-hacia-arriba (1).png" alt="">
                </button>
            </div>
        </div>

        <!-- Tabla -->
        <div class="overflow-x-auto p-4">
            <table class="min-w-full border border-gray-300 bg-white shadow-md">
                <thead class="bg-gray-100">
                    <tr>
                        <th class="border px-4 py-2">Ticker</th>
                        <th class="border px-4 py-2">Compañía</th>
                        <th class="border px-4 py-2">Rating (From → To)</th>
                        <th class="border px-4 py-2">Target (From → To)</th>
                        <th class="border px-4 py-2">Cambio (%)</th>
                        <th class="border px-4 py-2">Brokerage</th>
                        <th class="border px-4 py-2">Acción</th>
                        <th class="border px-4 py-2">Actualizado</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in carouselStore.items" :key="item.id" class="hover:bg-gray-50">
                        <td class="border px-4 py-2">{{ item.ticker }}</td>
                        <td class="border px-4 py-2">{{ item.company }}</td>
                        <td class="border px-4 py-2">{{ item.rating_from }} → {{ item.rating_to }}</td>
                        <td class="border px-4 py-2">{{ item.target_from }} → {{ item.target_to }}</td>
                        <td class="border px-4 py-2">
                            <span :class="getChangeColor(item)">
                                {{ getPercentChange(item) }}%
                            </span>
                        </td>
                        <td class="border px-4 py-2">{{ item.brokerage }}</td>
                        <td class="border px-4 py-2 flex items-center space-x-1">
                            <div :class="`w-2 h-2 rounded-full ${getActionColor(item.action)}`"></div>
                            <span>{{ item.action }}</span>
                        </td>
                        <td class="border px-4 py-2">{{ formatDate(item.time) }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useCarouselStore } from "@/stores/carouselStore";
import Card from "./Card.vue";

const carouselStore = useCarouselStore();

const formatDate = (isoDate: string) => {
    const date = new Date(isoDate);
    return date.toLocaleDateString("es-CO", {
        year: "numeric",
        month: "long",
        day: "numeric",
        timeZone: "America/Bogota"
    });
};

const getPercentChange = (item: any) => {
    const from = parseFloat(item.target_from.replace("$", ""));
    const to = parseFloat(item.target_to.replace("$", ""));
    if (isNaN(from) || from === 0) return "0.00";
    return ((to - from) / from * 100).toFixed(2);
};

const getChangeColor = (item: any) => {
    const change = parseFloat(getPercentChange(item));
    return change > 0 ? "text-green-500" : "text-red-500";
};

const getActionColor = (action: string) => {
    if (action.includes("upgraded") || action.includes("raised")) {
        return "bg-green-500"; 
    } else if (action.includes("lowered") || action.includes("downgraded")) {
        return "bg-red-500"; 
    } else {
        return "bg-gray-500"; 
    }
};
</script>
