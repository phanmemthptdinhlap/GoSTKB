const DataCard = {
  props: {
    index: {type: Number, default: 0},
    title: String,
    items: Array,
    click: Function,
  },
  template: `
    <div class="card" @click="clickHandler" style="margin: 5px; background-color: #17e6cd; border-radius: 5px;">
    <table style="width: auto;">
      <tr>
        <td>
          <div class="card-header" style="padding: 10px; text-align: center; width: 150px; height: 100%; background-color: #1755e6; color: white; border-radius: 5px;">
            <h3 class="card-title" >{{ title }}</h3>
          </div>
        </td>
        <td>
        <div class="card-body" style="padding: 5px; border-radius: 5px;">
        <table>
        <tr>
        <template v-for="(item, index) in items" :key="index">
          <td>
            <div class="card-item" style="padding: 5px; border-radius: 5px; border: 1px solid #ccc;">
              <h5 class="card-title" style="color: #1755e6;">{{ item.label }}</h5>
              <p class="card-text" style="color: #1755e6;">{{ item.value }}</p>
            </div>
          </td>
        </template>
        </tr>
        </table>
        </div>
        </td>
        </table>
    </div>
  `,
  setup(props) {
    const clickHandler = () => {
      props.click(props.index);
    };
    return {
      clickHandler,
    };
  }
}
