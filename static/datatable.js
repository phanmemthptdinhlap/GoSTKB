const DataTable={
  props:{
    labels: { type: object, required: true, default: () => ({}) },
    datas: { type: object, required: true, default: () => ({}) }
  },
  template: `
    <div class="panel">
      <table class="table">
        <thead>
          <tr>
            <template v-for="col in labels" :key="col.key" style="padding: 10px; text-align: center;">
              <th v-if="col.type !== 'cell'" style="padding: 10px; text-align: center;">
                {{ col.title }}
              </th>
              <template v-else>
                <th v-for="cell in col.cells" :key="col.key + '-' + cell.key" style="padding: 10px; text-align: center;">
                  {{ cell.title }}
                </th>
              </template>
            </template>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in localDatas" 
          :key="row.id || row.code" 
          style="border-bottom: 1px solid #eee;"
          @click="rowClick(row)">
            <template v-for="col in labels" :key="col.key">
              <td v-if="col.type !== 'cell'" style="padding: 8px; text-align: center;">
                {{ row[col.key] }}
              </td>
              <template v-else>
                <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" 
                style="padding: 8px; text-align: center; color: #d32f2f;">
                  {{ row[col.key][cell.key] }}
                </td>
              </template>
            </template> 
          </tr>
        </tbody>
      </table>
    </div>
     `,
  setup(props) {
    const localDatas = Vue.ref(props.datas);
    const rowClick = (row) => {
      alert(JSON.stringify(row));
    };
    return {
      localDatas,
      rowClick
    };
  } 
};
    // 1. Khởi t
