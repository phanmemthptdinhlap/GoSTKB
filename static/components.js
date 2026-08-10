const DataCard = {
  props: {
    index: {type: Number, default: 0},
    title: String,
    items: Array,
    click: Function,
  },
  template: `
    <div class="card" @click="clickHandler">
    <table>
      <tr>
        <td>
          <div class="card-header" style="padding: 5px;">
            <h3 class="card-title" >{{ title }}</h3>
          </div>
        </td>
        <template v-for="(item, index) in items" :key="index">
          <td>
            <table style="padding: 10px;">
              <tr><td>{{ item.label }}</td></tr>
              <tr><td>{{ item.value }}</td></tr>
            </table>
          </td>
        </template>
    </table>
    </div>
  `,
  setup(props) {
    const clickHandler = () => {
      props.click(index);
    };
    return {
      clickHandler,
    };
  }
}
