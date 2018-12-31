<template>
    <Card>
        <Form>
            <FormItem label="挖掘机" :label-width="80">
                <Select v-model="formItem.device_id">
                    <Option v-for="item in devices" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
             <FormItem label="油量(升)" :label-width="80">
                <InputNumber v-model="formItem.amount" style="width: 100%"></InputNumber>
            </FormItem>
             <FormItem label="油价(元/升)" :label-width="80">
                <InputNumber v-model="formItem.price" style="width: 100%"></InputNumber>
            </FormItem>
            <FormItem label="总价（元）" :label-width="80">
                <InputNumber v-model="formItem.total_price" style="width: 100%"></InputNumber>
            </FormItem>

            <FormItem label="日期时间" :label-width="80">
                 <DatePicker v-model="formItem.create_time" type="datetime" format="yyyy-MM-dd HH:mm" placeholder="选择开始日期和时间" style="width: 100%"></DatePicker>
            </FormItem>

            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    data() {
        return {
            formItem: {
                device_id: 0,
                amount:0,
                price:0,
                total_price:0,
                create_time:"",
            },
            devices:[],
        }
    },
    created() {
         this.$http.get("/device/list").then((data) => {
            this.devices = data.data.data;
        });
    },
    watch: {
      'formItem.amount':function(){
        this.formItem.total_price = this.formItem.amount*this.formItem.price;
      },
      'formItem.price':function(){
        this.formItem.total_price = this.formItem.amount*this.formItem.price;
      }
    },
    methods: {
        onSave() {
            if (this.formItem.device_id ==0) {
                this.$Message.error('请选择挖掘机');
                return;
            }
            if (this.formItem.price == 0) {
                this.$Message.error('请输入油价');
                return;
            }
            if (this.formItem.amount == "") {
                this.$Message.error('请输入加油量');
                return;
            }
            if (this.formItem.total_price == "") {
                this.$Message.error('请输入总价格');
                return;
            }
            this.$http.post('/gas/add', this.formItem).then((data) => {
                if (data.data.code == 0) {
                    this.$Message.success("挖掘机加油记录保存成功");
                    this.$router.push({ path: "/" });
                } else {
                    this.$Message.error("挖掘机加油记录保存失败");
                }
            });
        }
    }
}
</script>