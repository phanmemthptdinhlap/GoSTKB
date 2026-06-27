// GenericDatasheet.js
const DataSheet = {
  props: {
    labels: { type: Array, required: true },
    datas: { type: Array, required: true }
  },
  emits: ['sync'],
  // Khai báo giao diện trực tiếp trong chuỗi template
  template: `
    <div class="datasheet-containerx; display: flex; gap: 10px;">
      <div class="toolbar" style="margin-bottom: 10px;">
        <button @click="openAddModal" style="margin-right: 5px;">Thêm mới</button>
        <button @click="syncData">Đồng bộ Server</button>
      </div>

      <table border="1" style="width: 100%; border-collapse: collapse; border-spacing: 0; border-color: gray;">
        <thead>
          <tr>
            <th v-for="col in labels" :key="col.key" style="padding: 8px;">{{ col.title }}</th>
            <th style="padding: 8px;">Thao tác</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(row, index) in localDatas" :key="index">
            <td v-for="col in labels" :key="col.key" style="padding: 8px;">{{ row[col.key] }}</td>
            <td style="padding: 8px; text-align: center;">
              <button @click="openEditModal(row, index)" style="margin-right: 5px;">Sửa</button>
              <button @click="deleteRow(index)" style="color: red;">Xóa</button>
            </td>
          </tr>
          <tr v-if="localDatas.length === 0">
            <td :colspan="labels.length + 1" style="text-align: center; padding: 8px;">Chưa có dữ liệu</td>
          </tr>
        </tbody>
      </table>

      <div v-if="isModalOpen" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.5); display: flex; justify-content: center; align-items: center;">
        <div style="background: white; padding: 20px; border-radius: 5px; min-width: 300px;">
          <h3>{{ isEditing ? 'Sửa dữ liệu' : 'Thêm dữ liệu mới' }}</h3>
          <form @submit.prevent="saveModal">
            <div v-for="col in labels" :key="col.key" style="margin-bottom: 10px;">
              <label style="display: block; margin-bottom: 5px;">{{ col.title }}</label>
              <input v-model="formData[col.key]" :type="col.type || 'text'" required style="width: 100%; padding: 5px;" />
            </div>
            <div style="text-align: right; margin-top: 15px;">
              <button type="button" @click="closeModal" style="margin-right: 10px;">Hủy</button>
              <button type="submit">Lưu (Local)</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  `,
  setup(props, { emit }) {
    // Sử dụng Vue.xxx vì chúng ta đang gọi từ CDN global
    const localDatas = Vue.ref([...props.datas]);
    const isModalOpen = Vue.ref(false);
    const isEditing = Vue.ref(false);
    const editingIndex = Vue.ref(-1);
    const formData = Vue.ref({});

    // Theo dõi dữ liệu truyền vào để cập nhật
    Vue.watch(() => props.datas, (newVal) => {
      localDatas.value = [...newVal];
    }, { deep: true });

    const openAddModal = () => {
      formData.value = {};
      props.labels.forEach(col => formData.value[col.key] = '');
      isEditing.value = false;
      isModalOpen.value = true;
    };

    const openEditModal = (row, index) => {
      formData.value = { ...row };
      editingIndex.value = index;
      isEditing.value = true;
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
    };

    const saveModal = () => {
      if (isEditing.value) {
        localDatas.value[editingIndex.value] = { ...formData.value };
      } else {
        const newData = { ...formData.value, id: Date.now() }; // ID tạm thời
        localDatas.value.push(newData);
      }
      closeModal();
    };

    const deleteRow = (index) => {
      if (confirm('Bạn có chắc muốn xóa?')) {
        localDatas.value.splice(index, 1);
      }
    };

    const syncData = () => {
      emit('sync', localDatas.value);
    };

    return {
      localDatas, isModalOpen, isEditing, formData,
      openAddModal, openEditModal, closeModal, saveModal, deleteRow, syncData
    };
  }
};
