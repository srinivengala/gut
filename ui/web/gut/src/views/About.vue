<template>
  <div class="about">
    <h1>This is an about page</h1>

  <div class="mdl-grid">
    <!-- 12 desktop, 8 tablet, 4 phone -->
    <div class="mdl-cell mdl-cell--12-col mdl-cell--8-col-tablet mdl-cell--4-col-phone"><p>col s12</p></div>
    <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-tablet mdl-cell--4-col-phone"><p>col s12 m4 l2</p></div>
    <div class="mdl-cell mdl-cell--8-col mdl-cell--4-col-tablet mdl-cell--4-col-phone"><p>col s12 m4 l8</p></div>
    <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-tablet mdl-cell--4-col-phone"><p>col s12 m4 l2</p></div>
  </div>
  <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone"><p>col s12 m6 l3</p></div>
    <div class="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone"><p>col s12 m6 l3</p></div>
    <div class="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone"><p>col s12 m6 l3</p></div>
    <div class="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone"><p>col s12 m6 l3</p></div>
  </div>
  </div>
</template>
<style lang="stylus">
.mdl-cell
  background grey
.mdl-cell
  &.mdl-cell--3-col
    color white
    background blue
</style>