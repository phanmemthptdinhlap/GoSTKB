const DataTable = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
    cellsmap: { type: Object, required: false } // Bảng tra cứu dùng chung cho mọi loại đối tượng
  },
  emits: ['row-updated'],
  template: `
    <div class="panel">
      <!-- BẢNG DỮ LIỆU -->
      <table class="table" style="width: 100%; border-collapse: collapse;">
        <thead style="background: #f5f5f5;">
          <tr>
            <template v-for="col in labels" :key="col.key">
              <th v-if="col.type == 'text'" style="padding: 10px; border: 1px solid #ccc; text-align: center;">
                {{ col.title }}
              </th>
              <template v-else-if="col.type == 'group'">
                <th v-for="cell in col.cells" :key="col.key + '-' + cell.key" style="padding: 10px; border: 1px solid #ccc; text-align: center;">
                  {{ cell.title }}
                </th>
              </template>
            </template>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in localDatas" 
              :key="row.id || row.lop_id || Math.random()" 
              style="border-bottom: 1px solid #eee; cursor: pointer;"
              @click="openEditModal(row)">
            <template v-for="col in labels" :key="col.key">
              <td v-if="col.type == 'text'" style="padding: 8px; text-align: center; border: 1px solid #ccc;">
                {{ row[col.key] }}
              </td>
              <template v-else-if="col.type == 'group'">
                <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" 
                    style="padding: 8px; text-align: center; color: #d32f2f; border: 1px solid #ccc;">
                  {{ cellsmap[row[col.key]?.[cell.key]?.[col.subkey]] || '-' }}
                </td>
              </template>
            </template> 
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- MODAL CHỈNH SỬA TỔNG QUÁT -->
    <div v-if="isModalOpen" 
         style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.4); display: flex; justify-content: center; align-items: center; z-index: 1000;">
      <div style="background: white; padding: 25px; border-radius: 8px; min-width: 450px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
        <h3 style="margin-top: 0; border-bottom: 1px solid #eee; padding-bottom: 10px;">
          Điều chỉnh thông tin
        </h3>
        
        <form @submit.prevent="saveModal">
          <div v-for="cell in groupCol.cells" :key="cell.key" style="margin-bottom: 12px; display: flex; align-items: center;">
            <label style="width: 100px; font-weight: 500;">{{ cell.title }}</label>
            
            <!-- Binding động đa năng dựa trên groupCol.key và groupCol.subkey -->
            <select v-model.number="editingRow[groupCol.key][cell.key][groupCol.subkey]" class="form-select" style="flex: 1; padding: 6px; border: 1px solid #ccc; border-radius: 4px;">
              <option :value="null">-- Trống --</option>
              <option v-for="opt in optionsList" :key="opt.id" :value="opt.id">
                {{ opt.text }}
              </option>
            </select>
          </div>
          
          <div style="text-align: right; margin-top: 20px;">
            <button type="button" @click="closeModal" style="margin-right: 10px; padding: 6px 15px; border: 1px solid #ccc; background: white; cursor: pointer; border-radius: 4px;">Hủy</button>
            <button type="submit" style="padding: 6px 15px; background: #0d6efd; color: white; border: none; border-radius: 4px; cursor: pointer;">Lưu Thay Đổi</button>
          </div>
        </form>
      </div>
    </div>
  `,
  setup(props, { emit }) {
    const { ref, computed, watch } = Vue;
    
    const localDatas = ref(props.datas);
    const editingRow = ref(null);
    let originalRowRef = null;
    const isModalOpen = ref(false);
    
    watch(() => props.datas, (newVal) => {
      localDatas.value = newVal;
    }, { immediate: true, deep: true });

    // Trích xuất toàn bộ cấu hình của cột group (Đảm bảo không null)
    const groupCol = computed(() => {
      return props.labels.find(col => col.type === 'group') || { key: '', subkey: '', cells: [] };
    });

    // Chuyển đổi cellsmap thành mảng chuẩn để render thẻ Option
    const optionsList = computed(() => {
      const list = [];
      if (props.cellsmap) {
        for (const [id, text] of Object.entries(props.cellsmap)) {
          list.push({ id: Number(id), text: text });
        }
      }
      return list;
    });

    const openEditModal = (row) => {
      originalRowRef = row;
      editingRow.value = JSON.parse(JSON.stringify(row));
      
      const gCol = groupCol.value;
      if (gCol.key) {
        if (!editingRow.value[gCol.key]) {
          editingRow.value[gCol.key] = {};
        }

        // Dùng defaultObj từ config để khởi tạo trường thiếu, nếu không có fallback về null cho subkey
        gCol.cells.forEach(cell => {
          if (!editingRow.value[gCol.key][cell.key]) {
            editingRow.value[gCol.key][cell.key] = gCol.defaultObj 
              ? JSON.parse(JSON.stringify(gCol.defaultObj)) 
              : { [gCol.subkey]: null };
          }
        });
      }

      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
      editingRow.value = null;
      originalRowRef = null;
    };

    const saveModal = () => {
      Object.assign(originalRowRef, editingRow.value);
      emit('row-updated', originalRowRef);
      closeModal();
    };     
    
    return {
      localDatas,
      groupCol,
      optionsList,
      editingRow,
      isModalOpen,
      closeModal,
      saveModal,
      openEditModal
    };
  } 
};
