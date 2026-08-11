vue.component('checktable', {
  props: {
    title: String,
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
    theme: { 
      type: Object, 
      required: false,
      default: () => ({
        panel: 'check-panel',
        title: 'check-title',
        table: 'check-table',
        thead: 'check-thead',
        tr: 'check-tr',
        th: 'check-th',
        th_cell: 'check-th-cell',
        tbody: 'check-tbody',
        td: 'check-td',
        td_cell: 'check-td-cell',
        input: 'check-input',
        input_dirty: 'check-input-dirty',
        span: 'check-span'
      })
    },
  },
  template: `
  <div :class="theme.panel">
    <h3 :class="theme.title">{{ title }}</h3>
    <table :class="theme.table">
      <thead :class="theme.thead">
        <tr :class="theme.tr">
          <th :class="theme.th">...</th>
          <template v-for="col in labels.cols" :key="col.key">
            <th :class="theme.th">
              {{ col.text }}
            </th>
          </template>
          <th :class="theme.th">
            Chọn toàn bộ
          </th>
        </tr>
      </thead>
      <tbody :class="theme.tbody">
        <tr v-for="row in lables.rows" :key="row.key" :class="theme.tr">
            <td :class="theme.td">
              {{ row.text }}
            </td>
          <template v-for="cell in labels.cols" :key="cell.key">
            <td :class="theme.td_cell" style="text-align: center;">
              <input type="checkbox" v-model="ischecked(row.key, cell.key)" @change="checkRow(row.key, cell.key)">
            </td>
          </template>
            <td :class="theme.td">
              <input type="checkbox" v-model="isallchecked(row.key)" @change="checkAll(row.key)">
            </td>
        </tr>
      </tbody>
    </table>
    <div :class="theme.span">
      <span v-if="hasChanges">
        Đã có {{ changedPayload.length }} thay đổi
      </span>
      <span v-else>
        Không có thay đổi
      </span>
    </div>
  </div>
  `,
  setup(props) {
    const localDatas = Vue.ref([]);
    
    const initData = (newData) => {
      // SỬA LỖI: Deep Clone để tách biệt hoàn toàn bộ nhớ với props gốc
      const cloned = JSON.parse(JSON.stringify(newData));
      
      cloned.forEach(row => {
        props.labels.forEach(col => {
          if (col.type === 'cell' && row[col.key]) {
            col.cells.forEach(cellDef => {
              const cellData = row[col.key][cellDef.key];
              if (cellData) {
                // Lưu giá trị gốc để so sánh
                cellData._original = cellData[col.valuekey];
                cellData._isDirty = false;
              }
            });
          }
        });
      });
      localDatas.value = cloned;
    };

    // SỬA LỖI: Viết hoa 'Vue.watch' thay vì 'vue.watch'
    Vue.watch(
      () => props.datas,
      (newVal) => {
        if (newVal) {
          initData(newVal);
        }
      },
      { immediate: true, deep: true }
    );
    
    // 2. Hàm kiểm tra sự thay đổi của từng ô dữ liệu
    const checkDirtyState = (cellData, valuekey) => {
      cellData._isDirty = (cellData[valuekey] !== cellData._original);
    };

    // 3. Computed quét và trích xuất payload tự động
    const changedPayload = Vue.computed(() => {
      const payload = [];
      localDatas.value.forEach(row => {
        props.labels.forEach(col => {
          if (col.type === 'cell' && row[col.key]) {
            col.cells.forEach(cellDef => {
              const cellData = row[col.key][cellDef.key];
              
              if (cellData && cellData._isDirty) {
                // Đóng gói đối tượng theo đúng subkey và valuekey
                payload.push(cellData);
              }
            });
          }
        });
      });
      return payload;
    });

    const hasChanges = Vue.computed(() => changedPayload.value.length > 0);
   
    return {
      localDatas,
      checkDirtyState,
      changedPayload,
      hasChanges,
    };
  }
});


