const DataTable={
  props:{
    labels: { type: Object, required: true },
    datas: { type: Object, required: true },
    cellsmap: { type: Object, required: false }
  },
  template: `
    <div class="panel">
      <table class="table">
        <thead>
          <tr>
            <template v-for="col in labels" :key="col.key" style="padding: 10px; text-align: center;">
              <th v-if="col.type == 'text'" style="padding: 10px; text-align: center;">
                {{ col.title }}
              </th>
              <template v-else-if="col.type == 'group'">
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
              <td v-if="col.type == 'text'" style="padding: 8px; text-align: center;">
                {{ row[col.key] }}
              </td>
              <template v-else-if="col.type == 'group'">
                <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" 
                style="padding: 8px; text-align: center; color: #d32f2f;">
                  {{ cellsmap[row[col.key]?.[cell.key]?.[col.subkey]]||'-' }}
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
