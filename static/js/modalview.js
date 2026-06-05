const ModalView ={
  template: `
    <div class="modal" v-if="showModal" >
      <div class="modal-header">
        <div class="modal-title"> {{ modalTitle }} </div>
      </div>

      <div class="modal-body">
        <div class="modal-form">
          <div v-for="(label,index) in labels" :key="label.key" class="modal-item" >
            <label class="modal-item-label" style="display:inline-block; width: 80px;"> {{ label.label }}: </label>
            <input class="modal-item-input" type="text" v-model="data[label.key] ">
          </div>
        </div>
        <div class="modal-button">
          <button class="modal-button modal-button-save" @click="saveData" :disabled="isLoading">
            {{ isLoading ? 'Đang lưu ... ' : 'Lưu' }}
          </button>
          <button class="modal-button modal-button-cancel" @click="closeModal" :disabled="isLoading">
            Hủy
          </button>
        </div>
      </div>
      <div class="modal-footer">
        <div class = "modal-message">{{ message }}</div>
    </div>
  </div>
    `,
  props : {
    labels: { type: Array, default: () => [] },
  },
  emits : ['saveData'],

  setup(props, context){
    const {ref} = Vue;
    const {emit, expose} = context;
    const showModal = ref(false);
    const isEdit = ref(false);
    const isLoading = ref(false);
    const modalTitle = ref('Thêm mới');
    const data = ref({});
    const message = ref('');
      // Hàm xử lý dữ liệu
    const closeModal = () => {
      showModal.value = false;
      isEdit.value = false;
      data.value = {};
      };
    const openEditModal = (dataEdit,title) => {
      showModal.value = true;
      isEdit.value = true;
      modalTitle.value = title||'Chỉnh sửa';
      data.value = {...dataEdit};
      };
    const openAddModal = (title) => {
      showModal.value = true;
      isEdit.value = false;
      modalTitle.value = title||'Thêm mới';
      data.value = {};
      };
    const saveData = () => {
      isLoading.value = true;
      emit('saveData', data.value,(res) => {
        isLoading.value = false;
        if(res.success){
            closeModal();
          } else {
            message.value = res.message;
          }
        });
      };

    expose({
        openEditModal,
        openAddModal,
        closeModal,
      });
      return {
        showModal,
        isEdit,
        isLoading,
        modalTitle,
        data,
        message,
        saveData,
        closeModal,
      };
     
    }
};
