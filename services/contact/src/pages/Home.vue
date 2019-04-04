<template>
  <layout>
    <b-row>
      <b-col class="mb-4" cols="12" md="12" lg="4">
        <b-card class="text-center">
          <h1 class="display-4 text-info text-truncate w-100">1 023 456</h1>
          <div class="card-text text-secondary">
            <h3>Contacts</h3>
          </div>
        </b-card>
      </b-col>
      <b-col class="mb-4" cols="12" md="6" lg="4">
        <b-card class="text-center">
          <h1 class="display-4 text-success text-truncate w-100">0</h1>
          <div class="card-text text-secondary">
            <h3>News</h3>
          </div>
        </b-card>
      </b-col>
      <b-col class="mb-4" cols="12" md="6" lg="4">
        <b-card class="text-center">
          <h1 class="display-4 text-danger text-truncate w-100">12 456</h1>
          <div class="card-text text-secondary">
            <h3>Unsubscribers</h3>
          </div>
        </b-card>
      </b-col>
    </b-row>
  </layout>
</template>

<script>
export default {
  name: "home"
};
</script>
