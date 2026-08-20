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
          <template v-for="(ctext,ckey) in labels.cols" :key="ckey">
            <th :class="theme.th">
             <input 
                type="checkbox" 
                :checked="isCheckedAll(ckey)"
                @change="checkAll(ckey)"
                @change=""
                :class="theme.input"
              />
            {{ typeof ctext === 'object' ? ctext.text : ctext }}
            </th>
          </template>
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
              <input 
                type="checkbox" 
                :checked="isChecked(rkey,index)" 
                @change="checkRow(rkey,index)"
                :class=" changedMap[rkey+'_'+index] ? theme.input_dirty : theme.input"
              />
            </td>
          </template>
          </td>
        </tr>
      </tbody>
    </table>

  </div>
  `,
  setup(props) {
    const {computed, toRaw} = Vue;
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
    const buton_click=()=>{
      alert(JSON.stringify(changedMap.value,null,2))
    }

    // --- BỔ SUNG LỖI 2: Khai báo 4 hàm kiểm tra và xử lý checkbox ---

    const changedMap = computed(() => {
      const map = {};
      if (!localDatas.value||!props.datas) return map;
      Object.keys(localDatas.value).forEach(rkey => {
        const row = localDatas.value[rkey];
        const prow = props.datas[rkey];
        if (row && prow) {
          row.forEach((cell, cindex) => {
            if (cell !== prow[cindex]) {
              map[`${rkey}_${cindex}`] = true;
            }
          });
        }
      });
      return map;
    });

    const isCheckedAll = (rowKey) => {
      return !!localDatas.value[rowKey]?.every(cell => cell);
    };

    const checkAll = () => {
      localDatas.value.forEach(row => {
        row.forEach((cell, cindex) => {
          localDatas.value[rowKey][colKey] = !isChecked(rowKey, colKey);
        });
      });
    };

    const isChecked = (rowKey, colKey) => {
      return !!localDatas.value[rowKey]?.[colKey];
    };
    const checkRow = (rowKey, colKey) => {
      if (localDatas.value?.[rowKey]?.[colKey]!==undefined) {
        localDatas.value[rowKey][colKey] = !isChecked(rowKey, colKey);
      } else {
        console.log("Chưa có dữ liệu");
      }
      console.log(localDatas.value);
    };
    const hasChanges = computed(() => {
      return changedMapCount.value > 0;
    });
    const changedMapCount = computed(() => {
      if (!changedMap.value) return 0;
      return Object.keys(changedMap.value).length;
    });
    // Sửa lỗi 3: Sửa 'lables.rows' thành 'labels.rows'

    return {
      localDatas,
      isChecked, 
      checkRow,
      checkAll,
      isCheckedAll,
      changedMap,
      hasChanges,
      changedMapCount,
    };
  }
};
