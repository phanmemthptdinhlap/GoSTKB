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
          <template v-for="(ctext,cindex) in labels.cols" :key="cindex">
            <th :class="theme.th">
             <input 
                type="checkbox" 
                :checked="colCheckAll[cindex]"
                @change="checkAll(cindex)"
                :class="theme.input"
              />
            {{ typeof ctext === 'object' ? ctext.text : ctext }}
            </th>
          </template>
        </tr>
      </thead>
      <tbody :class="theme.tbody">
        <!-- Sửa lỗi 3: Sửa 'lables.rows' thành 'labels.rows' -->
        <tr v-for="row in labels.rows" :key="row" :class="theme.tr">
            <td :class="theme.td">
              {{ typeof row === 'object' ? row.text : row }}
            </td>
          <template v-for="(cell, index) in labels.cols" :key="index">
            <td :class="theme.td_cell" style="text-align: center;">
              <!-- Sửa lỗi 1: Thay v-model thành :checked -->
              <input 
                type="checkbox" 
                :checked="isChecked(row,index)"
                @change="checkRow(row,index)"
                :class=" (changedMap[row+'_'+index]!==undefined) ? theme.input_dirty : theme.input"
              />
            </td>
          </template>
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


    // --- BỔ SUNG LỖI 2: Khai báo 4 hàm kiểm tra và xử lý checkbox ---

    const changedMap = computed(() => {
      const map = {};
      if (!localDatas.value||!props.datas) return map;
      Object.keys(localDatas.value).forEach(rIndex => {
        const row = localDatas.value[rIndex];
        const prow = props.datas[rIndex];
        if (row && prow) {
          row.forEach((cell, cIndex) => {
            if (cell !== prow[cIndex]) {
              map[`${rIndex}_${cIndex}`] = row[cIndex];
            }else{
              delete map[`${rIndex}_${cIndex}`];
            }
          });
        }
      });
      return map;
    });

    const colCheckAll = computed(() => {
      const col={};
      if (!localDatas.value|| Object.keys(localDatas.value).length===0) return col;
      Object.keys(props.labels.cols).forEach((_, index) => {
        col[index] = Object.values(localDatas.value).every(row => {
          return !!row[index];
        });
      });
      return col;
    });

    const checkAll = (cIndex) => {
      if (!localDatas.value) {
        console.log("Chưa có dữ liệu checkAll");
        return;
      }
      const newState = !colCheckAll.value[cIndex];
      console.log(newState+" of State:"+cIndex);
      Object.keys(localDatas.value).forEach(rIndex => {
        if (localDatas.value[rIndex]!==undefined) {
          localDatas.value[rIndex][cIndex] = newState;
        }
        else {
          console.log("Chưa có dữ liệu "+rIndex);
        }
      });
    };

    const isChecked = (rKey, cIndex) => {
      return !!localDatas.value[rKey]?.[cIndex];
    };
    const checkRow = (rKey, cIndex) => {
      console.log(rKey,cIndex);
      if (localDatas.value?.[rKey]?.[cIndex]!==undefined) {
        localDatas.value[rKey][cIndex] = !isChecked(rKey, cIndex);
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
      colCheckAll,
      changedMap,
      hasChanges,
      changedMapCount,
    };
  }
};
