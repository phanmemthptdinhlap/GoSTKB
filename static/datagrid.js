// DataGrid.js
// DataGrid.js
const DataGrid = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
    theme: { 
      type: Object, 
      required: false,
      default: () => ({
        panel: 'grid-panel',
        table: 'grid-table',
        thead: 'grid-thead',
        tr: 'grid-tr',
        th: 'grid-th',
        th_cell: 'grid-th-cell',
        tbody: 'grid-tbody',
        td: 'grid-td',
        td_cell: 'grid-td-cell',
        input: 'grid-input',
        input_dirty: 'grid-input-dirty',
        span: 'grid-span'
      })
    },
  },
  template: `
  <div :class="theme.panel">
    <table :class="theme.table">
      <thead :class="theme.thead">
        <tr :class="theme.tr">
          <template v-for="col in labels" :key="col.key">
            <th v-if="col.type !== 'cell'" :class="theme.th">
              {{ col.title }}
            </th>
            <template v-else>
              <th v-for="cell in col.cells" :key="col.key + '-' + cell.key" :class="theme.th_cell">
                {{ cell.title }}
              </th>
            </template>
          </template>
        </tr>
      </thead>
      <tbody :class="theme.tbody">
        <tr v-for="(item, index) in localDatas" :key="item.lop_id || index" :class="theme.tr">
          
          <template v-for="col in labels" :key="col.key">
            <td v-if="col.type !== 'cell'" :class="theme.td">
              {{ item[col.key] }}
            </td>
            
            <template v-else>
              <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" :class="theme.td_cell">
                <input
                  v-if="item[col.key]?.[cell.key]"
                  type="number" 
                  v-model.number="item[col.key][cell.key][col.valuekey]"
                  @input="checkDirtyState(item[col.key][cell.key], col.valuekey)"
                  :class="item[col.key][cell.key]._isDirty ? theme.input_dirty : theme.input"
                  min="0"
                />
                <span v-else :class="theme.span"> - </span>
              </td>
            </template>
          </template>
          
        </tr>
      </tbody>
    </table>
  </div>
  `,
  // ... phần setup() giữ nguyên như cũ ...
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
};
