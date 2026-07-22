const DataTable = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
    cellsmap: { type: Object, required: false }
  },

  template: `
    <div class="panel">
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
          <tr v-for="(row,index) in localDatas" 
              :key="index" 
              style="border-bottom: 1px solid #eee; cursor: pointer;"
              @click="openEditModal(row)">
            <template v-for="col in labels" :key="col.key">
              <td v-if="col.type == 'text'" style="padding: 8px; text-align: center; border: 1px solid #ccc;">
                {{ row[col.key] }}
              </td>
              <template v-else-if="col.type == 'group'">
                <!-- Thêm dynamic style để bôi đỏ nếu hàm isDirty trả về true -->
                <td v-for="cell in col.cells" :key="col.key + '-' + cell.key" 
                    style="padding: 8px; text-align: center; border: 1px solid #ccc;"
                    :style="isDirty(index, col.key, cell.key, col.subkey) ? 'color: #d32f2f; font-weight: bold; background-color: #ffebee;' : 'color: inherit;'">
                  {{ cellsmap[row[col.key]?.[cell.key]?.[col.subkey]] || '-' }}
                </td>
              </template>
            </template> 
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- MODAL CHỈNH SỬA TỔNG QUÁT (Không đổi) -->
    <div v-if="isModalOpen" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.4); display: flex; justify-content: center; align-items: center; z-index: 1000;">
      <div style="background: white; padding: 25px; border-radius: 8px; min-width: 450px;">
        <h3 style="margin-top: 0; border-bottom: 1px solid #eee; padding-bottom: 10px;">Điều chỉnh thông tin</h3>
        <form @submit.prevent="saveModal">
          <div v-for="cell in groupCol.cells" :key="cell.key" style="margin-bottom: 12px; display: flex; align-items: center;">
            <label style="width: 100px; font-weight: 500;">{{ cell.title }}</label>
            <select v-model.number="editingRow[groupCol.key][cell.key][groupCol.subkey]" class="form-select" style="flex: 1; padding: 6px; border: 1px solid #ccc;">
              <option :value="null">-- Trống --</option>
              <option v-for="opt in optionsList" :key="opt.id" :value="opt.id">{{ opt.text }}</option>
            </select>
          </div>
          <div style="text-align: right; margin-top: 20px;">
            <button type="button" @click="closeModal" style="margin-right: 10px; padding: 6px 15px;">Hủy</button>
            <button type="submit" style="padding: 6px 15px; background: #0d6efd; color: white; border: none;">Lưu Tạm</button>
          </div>
        </form>
      </div>
    </div>
  `,
  setup(props) {
    const { ref, computed, watch } = Vue;
    
    const snapshotDatas = ref([]);
    const localDatas = ref([]);
    const editingRow = ref(null);
    let targetRowIndex = -1;
    const isModalOpen = ref(false);
    
    // Khởi tạo Snapshot và Local Working Copy
    watch(() => props.datas, (newVal) => {
      snapshotDatas.value = JSON.parse(JSON.stringify(newVal));
      localDatas.value = JSON.parse(JSON.stringify(newVal));
    }, { immediate: true, deep: true });

    const groupCol = computed(() => props.labels.find(col => col.type === 'group') || { key: '', subkey: '', cells: [] });
    
    const optionsList = computed(() => {
      const list = [];
      if (props.cellsmap) {
        for (const [id, text] of Object.entries(props.cellsmap)) {
          list.push({ id: Number(id), text: text });
        }
      }
      return list;
    });

    // So sánh dữ liệu hiện tại với Snapshot
    const isDirty = (currentRow, groupKey, cellKey, subKey) => {
      const idKey = currentRow.id !== undefined ? 'id' : 'id'; // Hỗ trợ linh hoạt khóa chính
      const originalRow = snapshotDatas.value.find(r => r[idKey] === currentRow[idKey]);
      
      if (!originalRow || !originalRow[groupKey] || !originalRow[groupKey][cellKey]) return false;
      
      const originalVal = originalRow[groupKey][cellKey][subKey];
      const currentVal = currentRow[groupKey]?.[cellKey]?.[subKey];
      
      return originalVal !== currentVal;
    };

    // Tự động quét và đóng gói các object thay đổi để cha gọi đồng bộ
    const changedPayload = computed(() => {
      const payload = [];
      const gCol = groupCol.value;
      if (!gCol.key) return payload;

      localDatas.value.forEach(row => {
        gCol.cells.forEach(cell => {
          if (isDirty(row, gCol.key, cell.key, gCol.subkey)) {
            // Đóng gói payload chuẩn cho API Golang
            payload.push({
              lop_id: row.id || row.lop_id,
              mon_id: parseInt(cell.key), // Môn học key thường là số
              giao_vien_id: row[gCol.key][cell.key][gCol.subkey]
            });
          }
        });
      });
      return payload;
    });

    const hasChanges = computed(() => changedPayload.value.length > 0);

    const openEditModal = (row) => {
      const idKey = row.id !== undefined ? 'id' : 'id';
      targetRowIndex = localDatas.value.findIndex(r => r[idKey] === row[idKey]);
      editingRow.value = JSON.parse(JSON.stringify(row));
      
      const gCol = groupCol.value;
      if (gCol.key) {
        if (!editingRow.value[gCol.key]) editingRow.value[gCol.key] = {};
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
      targetRowIndex = -1;
    };

    const saveModal = () => {
      // Chỉ lưu vào biến localDatas (Lưu tạm)
      if (targetRowIndex !== -1) {
        localDatas.value[targetRowIndex] = editingRow.value;
      }
      closeModal();
    };     
    
    // Expose để Component cha có thể truy cập bằng ref
    return {
      localDatas, groupCol, optionsList, editingRow, isModalOpen,
      closeModal, saveModal, openEditModal, isDirty,
      changedPayload, hasChanges 
    };
  } 
};
