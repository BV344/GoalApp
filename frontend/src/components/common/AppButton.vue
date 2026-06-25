<script setup lang="ts">
defineProps<{ label: string; variant?: 'primary' | 'secondary' }>()
defineEmits<{ click: [] }>()
</script>

<template>
  <button
    class="px-4 py-2 rounded font-medium transition"
    :class="variant === 'secondary' ? 'bg-gray-200 hover:bg-gray-300' : 'bg-blue-600 text-white hover:bg-blue-700'"
    @click="$emit('click')"
  >
    {{ label }}
  </button>
</template>
