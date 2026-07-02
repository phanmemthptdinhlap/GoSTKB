// GenericDatasheet.js
const DataSheet = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true }
  },
  emits: ['sync'],
  template: `
    <div class="datasheet-container" style="display: flex; flex-direction: column; gap: 10px;">
      
      <div class="toolbar" style="margin-bottom: 10px;">
        <table border="0" style="width: 100%; border-collapse: collapse; border-spacing: 0;">
          <tr>
            <td style="padding: 8px; text-align: left;">
              <button @click="openAddModal" style="margin-right: 5px;">Thêm mới</button>
              <button @click="syncData">Đồng bộ Server</button>
            </td>
            <td style="padding: 8px; text-align: right;">
              <input type="checkbox" v-model="isShowAll" id="chkShowAll" style="margin-right: 5px; cursor: pointer;" />
              <label for="chkShowAll" style="cursor: pointer; user-select: none;">Hiển thị trạng thái</label>
            </td>
          </tr>
        </table>
      </div>

      <table border="1" style="width: 100%; border-collapse: collapse; border-spacing: 0; border-color: gray;">
        <thead style="background-color: #f5f5f5;">
          <tr style="border-bottom: 1px solid #eee;">
            <th v-for="col in labels" :key="col.key" style="padding: 10px; text-align: left;">{{ col.title }}</th>
            <th v-if="isShowAll" style="padding: 10px; text-align: center; color: #d32f2f;">Trạng thái</th>
            <th style="padding: 10px; text-align: center; width: 120px;">Thao tác</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in displayedDatas" :key="row.id || row.code" style="border-bottom: 1px solid #eee;">
            <td v-for="col in labels" :key="col.key" style="padding: 8px;">
               <span :style="{ textDecoration: row.action === 'delete' ? 'line-through' : 'none', color: row.action === 'delete' ? '#9e9e9e' : 'inherit' }">
                 {{ row[col.key] }}
               </span>
            </td>
            <td v-if="isShowAll" style="padding: 8px; text-align: center; font-weight: bold; color: #d32f2f;">
               {{ row.action }}
            </td>
            <td style="padding: 8px; text-align: center; border-right: 1px solid #eee;">
              <template v-if="row.action !== 'xóa'">
                <button @click="openEditModal(row)" style="margin-right: 5px;">Sửa</button>
                <button @click="deleteRow(row)" style="color: red;">Xóa</button>
              </template>
              <template v-else>
                <button @click="restoreRow(row)" style="color: green;">Khôi phục</button>
              </template>
            </td>
          </tr>
          <tr v-if="displayedDatas.length === 0">
            <td :colspan="isShowAll ? labels.length + 2 : labels.length + 1" style="text-align: center; padding: 15px; color: gray;">
               Không có dữ liệu
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="isModalOpen" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.4); display: flex; justify-content: center; align-items: center;">
        <div style="background: white; padding: 25px; border-radius: 8px; min-width: 350px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
          <h3 style="margin-top: 0; border-bottom: 1px solid #eee; padding-bottom: 10px;">{{ isEditing ? 'Sửa thông tin' : 'Thêm dữ liệu mới' }}</h3>
          <form @submit.prevent="saveModal">
            <div v-for="col in labels" :key="col.key" style="margin-bottom: 12px;">
              <label style="display: block; margin-bottom: 5px; font-weight: 500;">{{ col.title }}</label>
              <input v-model="formData[col.key]" :type="col.type || 'text'" required style="width: 100%; padding: 8px; box-sizing: border-box; border: 1px solid #ccc; border-radius: 4px;" />
            </div>
            <div style="text-align: right; margin-top: 20px;">
              <button type="button" @click="closeModal" style="margin-right: 10px; padding: 6px 15px;">Hủy</button>
              <button type="submit" style="padding: 6px 15px; background-color: #1976d2; color: white; border: none; border-radius: 4px; cursor: pointer;">Lưu lại</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  `,
  setup(props, { emit }) {
    // 1. Khởi tạo mảng đúng chuẩn bằng .map()
    const localDatas = Vue.ref(props.datas.map(item => ({ ...item, action: ' ' })));
    
    // 2. Khai báo các biến trạng thái
    const isShowAll = Vue.ref(false); // Đã thêm
    const isModalOpen = Vue.ref(false);
    const isEditing = Vue.ref(false);
    const editingRowData = Vue.ref(null); // Dùng ref lưu đối tượng đang sửa thay vì index
    const formData = Vue.ref({});

    // Theo dõi dữ liệu từ props
    Vue.watch(() => props.datas, (newVal) => {
      localDatas.value = newVal.map(item => ({ ...item, action: ' ' }));
    }, { deep: true });

    // Computed hiển thị danh sách
    const displayedDatas = Vue.computed(() => {
      if (isShowAll.value) {
        return localDatas.value;
      }
      return localDatas.value.filter(row => row.action !== 'xóa');
    });

    const openAddModal = () => {
      formData.value = {};
      props.labels.forEach(col => formData.value[col.key] = '');
      isEditing.value = false;
      isModalOpen.value = true;
    };

    // 3. Xử lý Sửa bằng cách nhận thẳng object 'row'
    const openEditModal = (row) => {
      formData.value = { ...row };
      editingRowData.value = row;
      isEditing.value = true;
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
      editingRowData.value = null;
    };

    const saveModal = () => {
      if (isEditing.value && editingRowData.value) {
        // Cập nhật dữ liệu vào object đang tham chiếu
        Object.keys(formData.value).forEach(key => {
          editingRowData.value[key] = formData.value[key];
        });
        if (editingRowData.value.action === ' ') {
          editingRowData.value.action = 'sửa';
        }
      } else {
        // Đổi thành 'insert' để đồng nhất logic Backend Golang
        const newData = { ...formData.value, id: 0, action: 'thêm' }; 
        localDatas.value.push(newData);
      }
      closeModal();
    };

    // 4. Xóa chuẩn xác bằng object reference
    const deleteRow = (row) => {
      if (confirm('Bạn có chắc muốn xóa lớp học này?')) {
        const index = localDatas.value.indexOf(row);
        if (index > -1) {
          if (row.action === 'thêm') {
            localDatas.value.splice(index, 1); // Xóa hẳn nếu vừa thêm mới
          } else {
            row.action = 'xóa'; // Soft delete
          }
        }
      }
    };

    // Bổ sung tính năng khôi phục khi đang xem trạng thái
    const restoreRow = (row) => {
      row.action = ' ';
    };

    const syncData = () => {
      emit('sync', localDatas.value);
    };

    // 5. Bắt buộc phải return TẤT CẢ các biến/hàm được gọi trên template
    return {
      localDatas, isShowAll, displayedDatas, isModalOpen, isEditing, formData,
      openAddModal, openEditModal, closeModal, saveModal, deleteRow, restoreRow, syncData
    };
  }
};
