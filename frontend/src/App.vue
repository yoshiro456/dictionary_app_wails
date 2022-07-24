<script setup>
import { EMFindManyByWord } from '../wailsjs/go/main/App';

import { computed, ref } from '@vue/reactivity';
import debounce from 'lodash/debounce'
import { onMounted, onUnmounted, watch } from 'vue';

const iSearch = ref('')
const iSelectedIndex = ref(0)
const iSearchResult = ref([])
const refSearchBar = ref(null)
const oSelectedDictionary = computed(() => iSearchResult.value[iSelectedIndex.value] ?? null)
// TODO: 
// const oSelectedDictionaryHasSpeech = computed(() => oSelectedDictionary.value['word'])

const searchDictionary = () => {
  if (!iSearch.value) return
  EMFindManyByWord(iSearch.value).then((v) => {
    iSearchResult.value = v
  })
  // prevent invalid index

}
const throttleSearch = debounce(searchDictionary, 300)

watch(iSearch, throttleSearch)

const cleanSearchBar = () => iSearch.value = ''
const selectDictionary = (i) => iSelectedIndex.value = i
const focusSearchBar = () => refSearchBar.value.focus()

// move selection
const selectTopDictionary = () => {
  if (iSelectedIndex.value <= 0) return
  iSelectedIndex.value--
}
const selectBottomDictionary = () => {
  if (iSelectedIndex.value >= iSearchResult.value.length - 1) return
  iSelectedIndex.value++
}

const checkWordHasSpeech = (word) => {

}

const handleEsc = (e) => {
  if (e.key !== 'Escape') return
  e.preventDefault()
  cleanSearchBar()
  focusSearchBar()
}
const handleUp = (e) => {
  if (e.key !== 'ArrowUp') return
  e.preventDefault()
  selectTopDictionary()
  focusSearchBar()
}
const handleDown = (e) => {
  if (e.key !== 'ArrowDown') return
  e.preventDefault()
  selectBottomDictionary()
  focusSearchBar()
}
onMounted(() => {
  window.addEventListener('keydown', (e) => {
    handleEsc(e)
    handleUp(e)
    handleDown(e)
  })
})

onUnmounted(() => {
  window.removeEventListener('keydown', (e) => {
    handleEsc(e)
    handleUp(e)
    handleDown(e)
  })
})

</script>

<template>
  <div class="w-screen h-screen">
    <!-- container -->
    <div class="flex flex-col w-full h-full">
      <!-- top layout -->
      <div class="flex flex-shrink-0 w-full h-12 border-b-2 divide-x-2">
        <!-- search bar -->
        <div class="flex flex-grow-0 flex-shrink-0 w-52">
          <input ref="refSearchBar" @keydown.esc="handleEsc" @keydown.up="handleUp" @keydown.down="handleDown"
            class="flex flex-grow-0 w-full h-auto px-2 m-2 font-semibold border-2 border-black focus:outline-none"
            type="text" v-model="iSearch" />
        </div>
        <!-- Selected Result title -->
        <span class="flex items-center px-4 text-xl font-semibold">
          <div class="flex w-8">
            <!-- TODO: -->
            <template>
              <!-- v-if="checkWordHasSpeech(oSelectedDictionary['word'])" -->
              <!-- icon volume -->
              <div class="flex p-1 text-white bg-gray-600 rounded-md">
                <!-- @click="speechText()" -->
                <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
                </svg>
              </div>
            </template>
          </div>
          <!-- text -->
          <span v-if="oSelectedDictionary" class="pl-3" v-text="oSelectedDictionary['word']"></span>
        </span>
      </div>
      <!-- bottom layout -->
      <div class="flex flex-grow overflow-y-hidden divide-x-2">
        <!-- sidebar -->
        <div ref="refDictionaryResult" class="flex flex-col flex-shrink-0 h-full overflow-y-scroll bg-gray-50 w-52">
          <template v-for="(dictionary, index) in iSearchResult" :key="index">
            <span class="flex p-2 border-b-2" :class="{ 'bg-blue-500 text-white': iSelectedIndex === index }"
              v-text="dictionary['word']" @click="selectDictionary(index)"></span>
          </template>
        </div>
        <!-- definition -->
        <pre v-if="oSelectedDictionary" class="flex w-full h-full p-2 overflow-y-scroll"
          v-text="oSelectedDictionary['content'] ?? ''"></pre>
      </div>
    </div>
  </div>

</template>

<style>
</style>
