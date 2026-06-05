const TableView ={
  template: `
  <div class="tv-main">
      <div class="tv-header">
        <div class="tv-title"> {{ title }} </div>
        <div class="tv-subtitle"> {{ subtitle }} </div>
      </div>
      <div class="tv-body">
        <table class="tv-table">
          <thead class="tv-table-header">
            <tr>
              <th v-for="label in itemsLabel" :key="label.key" > {{ label.label }} </th>
              <th v-if="showAction">Thao tác</th>
            </tr>
          </thead>
          <tbody class="tv-table-body">
            <tr v-for="(item,index) in itemsData" :key="index">
              <td v-for="label in itemsLabel" :key="label.key" > {{ item[label.key] }} </td>
              <td v-if="showAction" class="tv-table-cell-edit-button" >
                <button class="tv-button tv-button-edit" @click="openEditModal(index)">
                  <span class="tv-button-text">Sửa</span>
                </button>
                <button class="tv-button tv-button-delete" @click="deleteItem(index)">
                  <span class="tv-button-text">Xóa</span>
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot class="tv-table-footer">
            <tr>
              <td colspan="100%" style="padding: 8px;">
                <div class="tv-pagination">
                  <span class="tv-pagination-text">Tổng số: {{ itemsData.length }} </span>
                  <button class="tv-pagination-button" @click="openAddModal">Thêm mới</button>
                </div>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>
      <div class="tv-footer">
        <div class="tv-footer-text"> {{ footer }} </div>
      </div>
    </div>
  <div class="modal" v-if="showModal" >
      <div class="modal-header">
        <div class="modal-title"> {{ isEdit? 'Sửa thông tin' : 'Thêm mới' }} {{ modalTitle }} </div>
      </div>

      <div class="modal-body">
        <div class="modal-form">
          <div v-for="(label,index) in itemsLabel" :key="label.key" class="modal-item" >
            <label class="modal-item-label" style="display:inline-block; width: 80px;"> {{ label.label }}: </label>
            <input class="modal-item-input" type="text" v-model="modalItem[label.key] ">
          </div>
        </div>
        <div class="modal-button">
          <button class="modal-button modal-button-save" @click="saveModal">Lưu</button>
          <button class="modal-button modal-button-cancel" @click="closeModal">Hủy</button>
        </div>
      </div>
    </div>
    `,
  props: {
    title: { type: String, default: 'Danh sách' },
    subtitle: { type: String, default: '' },
    footer: { type: String, default: '' },
    itemsLabel: { type: Array, default: () => [] },
    itemsData: { type: Array, default: () => [] },
    showAction: { type: Boolean, default: false },
  },
  data(){
    return {
      showModal: false,
      isEdit: false,
      modalItem: {},
      modalTitle: '',
    }
  },
  methods: {
    openEditModal(index){
      this.isEdit = true;
      this.modalItem = {...this.itemsData[index]};
      this.modalTitle = this.title;
      this.showModal = true;
    },
    openAddModal(){
      this.isEdit = false;
      this.modalItem = {};
      this.modalTitle = this.title;
      this.showModal = true;
    },
    closeModal(){
      this.showModal = false;
    },
    saveModal(){
      this.$emit('updateItem', this.modalItem);
      this.closeModal();
    },
    deleteItem(index){
      if(confirm('Bạn có chắc muốn xóa?'))
        this.$emit('deleteItem', index);
    }
  }
};
