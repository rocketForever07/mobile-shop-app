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
              v-for="pro in products"
            >
              <v-hover v-slot:default="{ hover }">
                <v-card class="mx-auto" color="grey lighten-4" max-width="600">
                  <v-img
                    class="white--text align-end"
                    :aspect-ratio="16 / 9"
                    height="200px"
                    :src="pro.src"
                  >
                    <v-card-title>{{ pro.type }} </v-card-title>
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
                        pro.name
                      }}</a>
                    </div>
                    <div>${{ pro.price }}</div>
                  </v-card-text>
                </v-card>
              </v-hover>
            </div>
          </div>
          <div class="text-center mt-12">
            <v-pagination v-model="page" :length="6"></v-pagination>
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
export default {
  data: () => ({
    range: [0, 10000],
    page: 1,
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
    products: [
      {
        id: 1,
        name: "BLACK TEE",
        type: "Jackets",
        price: "18.00",
        src: require("../assets/img/shop/1.jpg"),
      },
      {
        id: 2,
        name: "WHITE TEE",
        type: "Polo",
        price: "40.00",
        src: require("../assets/img/shop/2.jpg"),
      },
      {
        id: 3,
        name: "Zara limited...",
        type: "Denim",
        price: "25.00",
        src: require("../assets/img/shop/3.jpg"),
      },
      {
        id: 4,
        name: "SKULL TEE",
        type: "Jackets",
        price: "30.00",
        src: require("../assets/img/shop/4.jpg"),
      },
      {
        id: 5,
        name: "MANGO WINTER",
        type: "Sweaters",
        price: "50.00",
        src: require("../assets/img/shop/5.jpg"),
      },
      {
        id: 6,
        name: "SHIRT",
        type: "Denim",
        price: "34.00",
        src: require("../assets/img/shop/6.jpg"),
      },
      {
        id: 7,
        name: "TRUCKER JACKET",
        type: "Jackets",
        price: "38.00",
        src: require("../assets/img/shop/7.jpg"),
      },
      {
        id: 8,
        name: "COATS",
        type: "Jackets",
        price: "25.00",
        src: require("../assets/img/shop/8.jpg"),
      },
      {
        id: 9,
        name: "MANGO WINTER",
        type: "Sweaters",
        price: "50.00",
        src: require("../assets/img/shop/9.jpg"),
      },
      {
        id: 10,
        name: "SHIRT",
        type: "Denim",
        price: "34.00",
        src: require("../assets/img/shop/10.jpg"),
      },
      {
        id: 11,
        name: "TRUCKER JACKET",
        type: "Jackets",
        price: "38.00",
        src: require("../assets/img/shop/11.jpg"),
      },
      {
        id: 12,
        name: "COATS",
        type: "Jackets",
        price: "25.00",
        src: require("../assets/img/shop/12.jpg"),
      },
    ],
  }),
};
</script>
