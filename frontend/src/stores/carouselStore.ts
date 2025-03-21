import { defineStore } from "pinia";
import { ref, computed, onMounted, onUnmounted } from "vue";
import axios from "axios";

export const useCarouselStore = defineStore("carousel", () => {
    const items = ref([]);
    const currentIndex = ref(0);
    const cardsPerView = ref(3);
    const isLoading = ref(false);

    const API_URL = "http://localhost:3000/stocks";

    const fetchItems = async () => {
        if (isLoading.value) return;

        isLoading.value = true;

        try {
            const response = await axios.get(API_URL);
            if (Array.isArray(response.data)) {
                items.value = response.data.map(({ ticker, company, target_from, target_to, action, brokerage, rating_from, rating_to, time, recommended }) => ({
                    ticker,
                    company,
                    target_from,
                    target_to,
                    action,
                    brokerage,
                    rating_from,
                    rating_to,
                    time,
                    recommended
                }));
            } else {
                console.error("Respuesta inesperada de la API:", response.data);
            }
        } catch (error) {
            console.error("Error cargando datos:", error);
        } finally {
            isLoading.value = false;
        }
    };

    const updateCardsPerView = () => {
        const width = window.innerWidth;
        cardsPerView.value = width < 768 ? 1 : width < 1024 ? 2 : 3;
    };

    onMounted(() => {
        updateCardsPerView();
        fetchItems();
        window.addEventListener("resize", updateCardsPerView);
    });

    onUnmounted(() => {
        window.removeEventListener("resize", updateCardsPerView);
    });

    const nextSlide = () => {
        if (currentIndex.value + cardsPerView.value < items.value.length) {
            currentIndex.value++;
        }
    };

    const prevSlide = () => {
        if (currentIndex.value > 0) {
            currentIndex.value--;
        }
    };

    const slideStyle = computed(() => ({
        transform: `translateX(-${currentIndex.value * (100 / cardsPerView.value)}%)`,
    }));

    return { items, currentIndex, cardsPerView, nextSlide, prevSlide, isLoading, slideStyle };
});
