const DataTable = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true },
  },
  template: `
    <div class="panel">
      <table class="table" style="width: 100%; border-collapse: collapse;">
        <thead>
          <tr>
            <template v-for="col in labels" :key="col.id">
              <th v-if="col.type == 'text'" style="padding: 10px; border: 1px solid #ccc; text-align: center; background: #f8f9fa;">
                {{ col.title }}
              </th>
              <template v-else-if="col.type == 'group'">
                <th v-for="cell in col.children" :key="col.id + '-' + cell.id" style="padding: 10px; border: 1px solid #ccc; text-align: center; background: #f8f9fa;">
                  {{ cell.title }}
                </th>
              </template>
            </template>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(row, idx) in localDatas"
              :key="row.lop_id || idx"
              style="border-bottom: 1px solid #eee; cursor: pointer; transition: background 0.2s;"
              class="hover-row"
              @click="openModal(idx)">
            
            <template v-for="col in labels" :key="col.id">
              <td v-if="col.type == 'text'" style="padding: 8px; text-align: center; border: 1px solid #ccc;">
                {{ col.options[row[col.id]] || "-" }}
              </td>
              <template v-else-if="col.type == 'group'">
                <!-- Binding style động để bôi đỏ nếu dữ liệu bị thay đổi so với bản gốc -->
                <td v-for="cell in col.children" :key="col.id + '-' + cell.id"
                    style="padding: 8px; text-align: center; border: 1px solid #ccc;"
                    :style="isDirty(idx, col.id, cell.id) ? 'color: #d32f2f; font-weight: bold; background-color: #ffebee;' : 'color: inherit;'">
                  {{ col.options[row[col.id]?.[cell.id]] || '-' }}
                </td>
              </template>
            </template> 
          </tr>
        </tbody>
      </table>

      <!-- MODAL -->
      <div class="modal" v-if="modal_show" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.4); display: flex; justify-content: center; align-items: center; z-index: 1000;">
        <div style="background: white; padding: 20px; border-radius: 8px; min-width: 400px; box-shadow: 0 4px 12px rgba(0,0,0,0.15);">
          <div style="text-align: right; margin-top: -10px; cursor: pointer;">
            <span class="close" style="color: #666; font-size: 24px;" @click="closeModal">&times;</span>
          </div>
          
          <div class="modal-header" style="margin-bottom: 15px; border-bottom: 1px solid #eee; padding-bottom: 10px;">
            <h4 style="margin: 0; color: #333;">Điều chỉnh phân công: {{ labels[0]?.options[editingRow[labels[0].id]] }}</h4>
          </div>
          
          <div class="modal-body">
            <form @submit.prevent="saveModal">
              <table class="table" style="width: 100%; border-collapse: collapse;">
                <template v-for="col in labels" :key="col.id">
                  <tr v-if="col.type == 'text'">
                    <th style="padding: 10px 0; text-align: left; width: 120px;">{{ col.title }}</th>
                    <td style="padding: 10px 0; text-align: left; font-weight: bold;">
                      {{ col.options[editingRow[col.id]] || "-" }}
                    </td>
                  </tr>
                  <template v-else-if="col.type == 'group'">
                    <tr v-for="cell in col.children" :key="col.id + '-' + cell.id">
                      <th style="padding: 10px 0; text-align: left;">{{ cell.title }}</th>
                      <td style="padding: 10px 0; text-align: left;">
                        <!-- v-model trỏ vào editingRow (bản sao tạm) thay vì localDatas -->
                        <select v-model="editingRow[col.id][cell.id]" style="width: 100%; padding: 8px; box-sizing: border-box; border: 1px solid #ccc; border-radius: 4px; outline: none;">
                            <option disabled value="">-- Chọn {{ cell.title }} --</option>   
                            <option v-for="(value, id) in col.options" :key="id" :value="id">
                              {{ value }}
                            </option>
                        </select> 
                      </td>
                    </tr>
                  </template>
                </template>
              </table>
              <div class="modal-footer" style="text-align: right; margin-top: 20px;">
                <button type="button" @click="closeModal" style="padding: 8px 16px; margin-right: 10px; border: 1px solid #ccc; background: white; border-radius: 4px; cursor: pointer;">Đóng</button>
                <button type="submit" style="padding: 8px 16px; background: #0d6efd; color: white; border: none; border-radius: 4px; cursor: pointer;">Xác nhận Lưu</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  `,
  setup(props) {
    const { ref, watch, computed } = Vue;
    
    const snapshotDatas = ref([]);
    const localDatas = ref([]);
    const editingRow = ref(null);
    const index = ref(-1);
    const modal_show = ref(false);
    
    // Tạo Deep Clone độc lập cho snapshot và local state
    watch(() => props.datas, (newVal) => {
      snapshotDatas.value = JSON.parse(JSON.stringify(newVal));
      localDatas.value = JSON.parse(JSON.stringify(newVal));
    }, { immediate: true, deep: true });

    const openModal = (idx) => {
      index.value = idx;
      // Khởi tạo bản sao tạm thời để làm việc trong Form
      editingRow.value = JSON.parse(JSON.stringify(localDatas.value[idx]));
      
      // Xử lý an toàn: Khởi tạo Object nếu chưa tồn tại dữ liệu môn học
      props.labels.forEach(col => {
        if (col.type === 'group' && !editingRow.value[col.id]) {
          editingRow.value[col.id] = {};
        }
      });
      
      modal_show.value = true;
    };

    const closeModal = () => {
      modal_show.value = false;
      editingRow.value = null;
    };

    const saveModal = () => {
      // Đẩy dữ liệu từ bản sao tạm vào lưới hiển thị chính
      localDatas.value[index.value] = JSON.parse(JSON.stringify(editingRow.value));
      closeModal();
    };

    // Thuật toán kiểm tra sự thay đổi O(1)
    const isDirty = (idx, colId, cellId) => {
      const original = snapshotDatas.value[idx]?.[colId]?.[cellId];
      const current = localDatas.value[idx]?.[colId]?.[cellId];
      // Nếu giá trị undefined/null quy về rỗng để tránh lỗi so sánh type
      return (original || "") !== (current || "");
    };

    // Đóng gói mảng dữ liệu đã thay đổi để API Server xử lý Upsert
    const changedPayload = computed(() => {
      const payload = [];
      const groupCol = props.labels.find(c => c.type === 'group');
      if (!groupCol) return payload;

      localDatas.value.forEach((row, idx) => {
        groupCol.children.forEach(cell => {
          if (isDirty(idx, groupCol.id, cell.id)) {
            payload.push({
              lop_id: row.lop_id,
              mon_id: parseInt(cell.id),
              giao_vien_id: parseInt(row[groupCol.id][cell.id])
            });
          }
        });
      });
      return payload;
    });

    const hasChanges = computed(() => changedPayload.value.length > 0);

    return {
      localDatas,
      editingRow,
      index,
      modal_show,
      openModal,
      closeModal,
      saveModal,
      isDirty,
      changedPayload,
      hasChanges
    };
  } 
};
