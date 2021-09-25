<template>
  <div>
    <v-container>
      <div class="row">
        <div class="col-md-3 col-sm-3 col-xs-12">
          <v-card outlined>
            <v-card-title>Manufacturer</v-card-title>
            <v-divider></v-divider>
            <template>
              <v-treeview
                :items="items"
                :open="[1]"
                :active="[5]"
                :selected-color="'#fff'"
                activatable
                open-on-click
                dense
              ></v-treeview>
            </template>
            <v-divider></v-divider>
            <v-card-title>Price</v-card-title>
            <v-range-slider
              v-model="range"
              :max="max"
              :min="min"
              :height="10"
              class="align-center"
              dense
            ></v-range-slider>
            <v-row class="pa-2" dense>
              <v-col cols="12" sm="5">
                <v-text-field
                  :value="range[0]"
                  label="Min"
                  outlined
                  dense
                  @change="$set(range, 0, $event)"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="2">
                <p class="pt-2 text-center">TO</p>
              </v-col>
              <v-col cols="12" sm="5">
                <v-text-field
                  :value="range[1]"
                  label="Max"
                  outlined
                  dense
                  @change="$set(range, 1, $event)"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-divider></v-divider>

            <v-divider></v-divider>
            <v-card-title class="pb-0">COLOR</v-card-title>
            <v-container class="pt-0" fluid>
              <v-checkbox label="BLACK" hide-details dense></v-checkbox>
              <v-checkbox label="WHITE" hide-details dense></v-checkbox>
              <v-checkbox label="RED" hide-details dense></v-checkbox>
              <v-checkbox label="GOLD" hide-details dense></v-checkbox>
              <v-checkbox label="GREEN" hide-details dense></v-checkbox>
              <v-checkbox label="MATTE BLACK" hide-details dense></v-checkbox>
            </v-container>
          </v-card>
        </div>
        <div class="col-md-9 col-sm-9 col-xs-12">
          <v-row dense>
            <v-col cols="12" sm="8" class="pl-6 pt-6">
              <small>Showing 1-12 of 200 products</small>
            </v-col>
          </v-row>

          <v-divider></v-divider>

          <div class="row text-center">
            <div
              class="col-md-3 col-sm-6 col-xs-12"
              :key="pro.id"
              v-for="pro in paginate"
            >
              <v-hover v-slot:default="{ hover }">
                <v-card class="mx-auto" color="grey lighten-4" max-width="600">
                  <v-img
                    class="white--text align-end"
                    :aspect-ratio="16 / 9"
                    height="200px"
                    :src="pro.Img"
                  >
                    <v-expand-transition>
                      <div
                        v-if="hover"
                        class="
                          d-flex
                          transition-fast-in-fast-out
                          white
                          darken-2
                          v-card--reveal
                          display-3
                          white--text
                        "
                        style="height: 100%"
                      >
                        <v-btn v-if="hover" href="/product" class="" outlined
                          >VIEW</v-btn
                        >
                      </div>
                    </v-expand-transition>
                  </v-img>
                  <v-card-text class="text--primary">
                    <div>
                      <a href="/product" style="text-decoration: none">{{
                        pro.Name
                      }}</a>
                    </div>
                    <div>${{ pro.Price }}</div>
                  </v-card-text>
                </v-card>
              </v-hover>
            </div>
          </div>
          <div class="text-center mt-12">
            <v-pagination
              :length="totalPages"
              circle
              @input="setPage"
              v-model="currentPage"
            >
            </v-pagination>
          </div>
        </div>
      </div>
    </v-container>
  </div>
</template>
<style>
.v-card--reveal {
  align-items: center;
  bottom: 0;
  justify-content: center;
  opacity: 0.8;
  position: absolute;
  width: 100%;
}
</style>
<script>
import axios from "axios";
export default {
  data: () => ({
    range: [0, 10000],
    // page: 2,
    breadcrums: [
      {
        text: "Home",
        disabled: false,
        href: "breadcrumbs_home",
      },
      {
        text: "Clothing",
        disabled: false,
        href: "breadcrumbs_clothing",
      },
      {
        text: "T-Shirts",
        disabled: true,
        href: "breadcrumbs_shirts",
      },
    ],
    min: 0,
    max: 10000,
    items: [
      {
        id: 2,
        name: "APPLE",
        children: [
          { id: 2, name: "IPHONE 13" },
          { id: 3, name: "IPHONE 12" },
          { id: 4, name: "IPHONE 11" },
          { id: 5, name: "IPHONE X" },
          { id: 6, name: "IPHONE 8" },
          { id: 7, name: "IPHONE 7" },
        ],
      },
      {
        id: 1,
        name: "SAMSUNG",
        children: [
          { id: 8, name: "GALAXY Z FOLD " },
          { id: 9, name: "GALAXY NOTE" },
          { id: 10, name: "GALAXY S" },
          { id: 11, name: "GALAXY A" },
        ],
      },
      {
        id: 3,
        name: "OPPO",
        children: [
          { id: 12, name: "OPPO FIND" },
          { id: 13, name: "OPPO RENO" },
          { id: 14, name: "OPPO A" },
        ],
      },
    ],
    products: [],

    currentPage: 1,
    itemsPerPage: 12,
    resultCount: 0,
  }),
  created() {
    axios
      .get("http://localhost:9999/api/product")
      .then((res) => {
        console.log(res.data);
        this.products = res.data;
      })
      .catch((error) => {
        console.log(error);
      });
  },
  methods: {
    setPage: function (pageNumber) {
      this.currentPage = pageNumber;
    },
  },
  computed: {
    totalPages: function () {
      this.resultCount = this.products.length;
      if (this.resultCount == 0) {
        return 1;
      } else {
        return Math.ceil(this.resultCount / this.itemsPerPage);
      }
    },
    paginate: function () {
      if (!this.products || this.products.length != this.products.length) {
        return;
      }
      this.resultCount = this.products.length;
      if (this.currentPage >= this.totalPages) {
        this.currentPage = this.totalPages;
      }
      var index = this.currentPage * this.itemsPerPage - this.itemsPerPage;
      return this.products.slice(index, index + this.itemsPerPage);
    },
  },
};
</script>
