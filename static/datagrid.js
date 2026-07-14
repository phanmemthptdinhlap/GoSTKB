// DataGrid.js
// DataGrid.js
const DataGrid = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
  },
  template: `
  <div>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <template v-for="col in labels" :key="col.key">
            <th v-if="col.type !== 'cell'" :class="col.header">
              {{ col.title }}
            </th>
            <template v-else>
              <th v-for="cell in col.cells" :key="col.key + '-' + cell.key" :class="col.header">
                {{ cell.title }}
              </th>
            </template>
          </template>
        </tr>
      </thead>
      <tbody>
        <!-- Đổi item.key thành item.lop_id hoặc lấy index để tránh warning -->
        <tr v-for="(item, index) in localDatas" :key="item.lop_id || index">
          
          <template v-for="col in labels" :key="col.key">
            
            <td v-if="col.type !== 'cell'" :class="col.header">
              {{ item[col.key] }}
            </td>
            
            <template v-else>
              <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" :class="col.header">
                <input
                  v-if="item[col.key]?.[cell.key]"
                  type="number" 
                  v-model.number="item[col.key][cell.key][col.valuekey]"
                  @input="checkDirtyState(item[col.key][cell.key], col.valuekey)"
                  class="form-control form-control-sm mx-auto"
                  :class="{'border-danger bg-danger-subtle': item[col.key][cell.key]._isDirty}"
                  style="width: 30px; text-align: center;"
                  min="0"
                />
                <span v-else class="text-muted">-</span>
              </td>
            </template>
            
          </template>
          
        </tr>
      </tbody>
    </table>
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
};
