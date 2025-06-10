<script setup>
import Chat from './components/Chat.vue'
</script>

<template>
  <header>
    <h1>汇编助手演示</h1>
  </header>
  <main>
    <Chat />
  </main>
</template>

<style>
html, body {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  background: linear-gradient(to right, #001e82, #9016c5);
  /* 或你喜欢的渐变色 */
  overflow-x: hidden;
}

#app {
  min-height: 100vh;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: linear-gradient(to right, #001e82, #9016c5);
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
}

header {
  text-align: center;
  margin: 0 0 30px 0;
  padding: 32px 0 0 0;
}

header h1 {
  color: #fff;
  font-size: 2.4em;
  letter-spacing: 2px;
  font-weight: 700;
  text-shadow: 0 2px 16px rgba(0,0,0,0.18);
}

main {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>