const CheckTable = {
  props: {
    title: String,
    labels: { type: Object, required: true }, // Sửa: Sửa kiểu dữ liệu thành Object ({ cols: [], rows: [] })
    datas: { type: Object, required: true },
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
        input_dirty: 'background-color: red;',
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
          <template v-for="(ctext,ckey) in labels.cols" :key="ckey">
            <th :class="theme.th">
              {{ typeof ctext === 'object' ? ctext.text : ctext }}
            </th>
          </template>
          <th :class="theme.th">
            Chọn toàn bộ
          </th>
        </tr>
      </thead>
      <tbody :class="theme.tbody">
        <!-- Sửa lỗi 3: Sửa 'lables.rows' thành 'labels.rows' -->
        <tr v-for="(rtext, rkey) in labels.rows" :key="rkey" :class="theme.tr">
            <td :class="theme.td">
              {{ typeof rtext === 'object' ? rtext.text : rtext }}
            </td>
          <template v-for="(cell, ckey, index) in labels.cols" :key="ckey">
            <td :class="theme.td_cell" style="text-align: center;">
              <!-- Sửa lỗi 1: Thay v-model thành :checked -->
              <label>{{rkey}} - {{index}}</label>
              <input 
                type="checkbox" 
                :checked="isChecked(rkey, index)"
                @change="checkRow(rkey, index)"
                :style="[theme.input, isChanged(rkey, index) ? theme.input_dirty : '']"
              >
            </td>
          </template>
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
    const {ref,watch,computed, onMounted, toRaw} = Vue;
    const localDatas = Vue.ref([]);
    
    const initData = (newData) => {
      localDatas.value = structuredClone(toRaw(newData));
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
    const isChanged = (rowKey, colKey) => {
      const current = !!(localDatas.value?.[rowKey]?.[colKey]);
      const original = !!(props.datas?.[rowKey]?.[colKey]);
      return current !== original;
    };
    const isChecked = (rowKey, colKey) => {
      return !!localDatas.value[rowKey]?.[colKey];
    };
    const checkRow = (rowKey, colKey) => {
      if (localDatas.value?.[rowKey]?.[colKey]) {
        localDatas.value[rowKey][colKey] = !isChecked(rowKey, colKey);
      }
    };
 
    return {
      localDatas,
      isChecked,
      isChanged,
      checkRow,
    };
  }
};
