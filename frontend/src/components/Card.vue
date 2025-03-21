<template>
    <div class="grid gap-4 py-8 justify-center ">
        <div>
            <div
                class="grid relative w-60 pt-14 p-6 bg-white border border-gray-200 rounded-lg shadow-sm dark:bg-gray-800 dark:border-gray-700 bg-gradient-to-b from-withe-50 to-green-50">
                <span
                    class="w-28 text-center absolute top-1 right-1  rounded-full bg-green-200 p-0.5 text-xs font-extralight px-4 py-1 justify-end">
                    Recomendation
                </span>

                <h3 class=" text-center mb-2 text-4xl font-bold tracking-tight text-gray-900 dark:text-white">
                    {{ item.ticker }}</h3>
                <span class="abosolute top-2 justify-self-center text-xs text-gray-700">{{ item.company }}</span>
                <div class="grid mt-6 grid-cols-2 ">
                    <div>
                        <span class="block top-2 text-xs  text-gray-700 font-extralight">To</span>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">
                            {{ item.rating_to }}</p>
                        <span class="block top-2 text-xs  text-gray-700 font-extralight">From</span>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">
                            {{ item.rating_from }}</p>
                    </div>
                    <div>
                        <span class="block top-2 text-xs text-right text-gray-700 font-extralight">To</span>
                        <span class="block top-2 text-xs text-right text-gray-700 font-extralight">
                            {{ percentChange >= 0 ? '+' : '' }}{{ percentChange.toFixed(2) }}%
                        </span>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 text-end">
                            {{ item.target_to }}</p>
                        <span class="block top-2  text-xs text-right  text-gray-700 font-extralight">From</span>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 text-end">
                            {{ item.target_from }}</p>
                    </div>

                </div>
                <div class="border-t-2 border-dashed border-gray-500 my-4"></div>

                <div>
                    <span class="block top-2 text-xs text-right text-black-700 font-extralight">brokerage</span>
                    <span class="block top-2 text-xs text-right text-gray-700 font-extralight">{{ item.brokerage }}</span>
                </div>
                <div class="mt-4">
                    <div class="grid grid-cols-2 ">
                        <div>
                            <span class="block top-2 text-xs text-left text-gray-700 font-extralight">action</span>
                            <div class="flex items-center space-x-1">
                                <div :class="`w-2 h-2 rounded-full ${getActionColor(item.action)}`"></div>
                                <span class="block top-2 text-xs text-left text-black-700 font-extralight truncate max-w-[100px]">{{item.action}}</span>
                            </div>

                        </div>
                        <div>
                            <span class="block top-2 text-xs text-right text-black-700 font-extralight">updated</span>
                            <span class="block top-2 text-xs text-right text-gray-700 font-extralight">{{  formatDate(item.time) }}</span>
                        </div>
                    </div>
                </div>

            </div>


        </div>


    </div>

</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    item: {
        ticker: string;
        target_from: string;
        target_to: string;
        company: string;
        action: string;
        brokerage: string;
        rating_from: string;
        rating_to: string;
        time: string;
    };
}>();

const formatDate = (isoDate: string) => {
  const date = new Date(isoDate);
  return date.toLocaleDateString("es-CO", {
    year: "numeric",
    month: "long",
    day: "numeric",
    timeZone: "America/Bogota"
  });
};




const percentChange = computed(() => {
    const from = parseFloat(props.item.target_from.replace("$", ""));
    const to = parseFloat(props.item.target_to.replace("$", ""));
    if (isNaN(from) || from === 0) return 0;
    return ((to - from) / from) * 100;
});

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