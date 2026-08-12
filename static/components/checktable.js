const CheckTable = {
  props: {
    title: String,
    labels: { type: Object, required: true }, // Sửa: Sửa kiểu dữ liệu thành Object ({ cols: [], rows: [] })
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
          <th :class="theme.th">Mục</th>
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
        <!-- Sửa lỗi 3: Sửa 'lables.rows' thành 'labels.rows' -->
        <tr v-for="row in labels.rows" :key="row.key" :class="theme.tr">
            <td :class="theme.td">
              {{ row.text }}
            </td>
          <template v-for="cell in labels.cols" :key="cell.key">
            <td :class="theme.td_cell" style="text-align: center;">
              <!-- Sửa lỗi 1: Thay v-model thành :checked -->
              <input 
                type="checkbox" 
                :checked="ischecked(row.key, cell.key)" 
                @change="checkRow(row.key, cell.key, $event)"
              >
            </td>
          </template>
            <td :class="theme.td" style="text-align: center;">
              <!-- Sửa lỗi 1: Thay v-model thành :checked -->
              <input 
                type="checkbox" 
                :checked="isallchecked(row.key)" 
                @change="checkAll(row.key, $event)"
              >
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
      const cloned = JSON.parse(JSON.stringify(newData));
      
      cloned.forEach(row => {
        // Sửa lỗi 4: Duyệt qua props.labels.cols
        if (props.labels && props.labels.cols) {
          props.labels.cols.forEach(col => {
            if (row[col.key]) {
              const cellData = row[col.key];
              cellData._original = cellData.checked || false;
              cellData._isDirty = false;
            }
          });
        }
      });
      localDatas.value = cloned;
    };

    Vue.watch(
      () => props.datas,
      (newVal) => {
        if (newVal) {
          initData(newVal);
        }
      },
      { immediate: true, deep: true }
    );

    // --- BỔ SUNG LỖI 2: Khai báo 4 hàm kiểm tra và xử lý checkbox ---

    // 1. Kiểm tra ô đơn lẻ có được chọn không
    const ischecked = (rowKey, colKey) => {
      return localDatas.value.find(r => r[0] === rowKey && r[1] === colKey)|| false;
    };

    // 2. Xử lý khi tick/untick ô đơn lẻ
    const checkRow = (rowKey, colKey, event) => {
      const isChecked = event.target.checked;
      const row = localDatas.value.find(r => r.key === rowKey);
      if (row && row[colKey]) {
        row[colKey].checked = isChecked;
        row[colKey]._isDirty = (row[colKey].checked !== row[colKey]._original);
      }
    };

    // 3. Kiểm tra xem toàn bộ các cột trong 1 dòng có được chọn hết không
    const isallchecked = (rowKey) => {
      const row = localDatas.value.find(r => r.key === rowKey);
      if (!row || !props.labels.cols) return false;
      return props.labels.cols.every(col => row[col.key] && row[col.key].checked);
    };

    // 4. Xử lý khi tick nút "Chọn toàn bộ" của 1 dòng
    const checkAll = (rowKey, event) => {
      const isChecked = event.target.checked;
      const row = localDatas.value.find(r => r.key === rowKey);
      if (row && props.labels.cols) {
        props.labels.cols.forEach(col => {
          if (row[col.key]) {
            row[col.key].checked = isChecked;
            row[col.key]._isDirty = (row[col.key].checked !== row[col.key]._original);
          }
        });
      }
    };

    // Quét và lọc danh sách các ô đã bị thay đổi
    const changedPayload = Vue.computed(() => {
      const payload = [];
      localDatas.value.forEach(row => {
        if (props.labels && props.labels.cols) {
          props.labels.cols.forEach(col => {
            const cellData = row[col.key];
            if (cellData && cellData._isDirty) {
              payload.push(cellData);
            }
          });
        }
      });
      return payload;
    });

    const hasChanges = Vue.computed(() => changedPayload.value.length > 0);
   
    return {
      localDatas,
      changedPayload,
      hasChanges,
      ischecked,
      checkRow,
      isallchecked,
      checkAll
    };
  }
};
