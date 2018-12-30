<template>
    <Card>
        <Form>
            <FormItem label="挖掘机" :label-width="80">
                <Select v-model="formItem.device_id">
                    <Option v-for="item in devices" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
            <FormItem label="司机" :label-width="80">
                <Select v-model="formItem.driver_id">
                    <Option v-for="item in drivers" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
             <FormItem label="客户" :label-width="80">
                <Input v-model="formItem.customer" placeholder="请输入客户名称..." />
            </FormItem>
             <FormItem label="地点" :label-width="80">
                <Input v-model="formItem.workplace" placeholder="请输入工作地点..." type="textarea" :rows="2"/>
            </FormItem>
             <FormItem label="备注" :label-width="80">
                <Input v-model="formItem.remarks" placeholder="请输入备注..." type="textarea" :rows="2"/>
            </FormItem>

            <FormItem label="开始时间" :label-width="80">
                 <DatePicker v-model="formItem.start_time" type="datetime" format="yyyy-MM-dd HH:mm" placeholder="选择开始日期和时间" style="width: 100%"></DatePicker>
            </FormItem>

            <FormItem label="结束时间" :label-width="80">
                <DatePicker v-model="formItem.end_time"  type="datetime" format="yyyy-MM-dd HH:mm" placeholder="选择结束日期和时间" style="width: 100%"></DatePicker>
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
                driver_id: 0,
                customer:"",
                workplace:"",
                remarks:"",
                start_time:"",
                end_time:"",
            },
            drivers: [],
            devices:[],
        }
    },
    created() {
        this.$http.get("/driver/list").then((data) => {
            this.drivers = data.data.data;
        });
         this.$http.get("/device/list").then((data) => {
            this.devices = data.data.data;
        });
    },
    methods: {
        onSave() {
            if (this.formItem.device_id ==0) {
                this.$Message.error('请选择挖掘机');
                return;
            }
            if (this.formItem.driver_id == 0) {
                this.$Message.error('请选择司机');
                return;
            }
            if (this.formItem.customer == "") {
                this.$Message.error('请输入客户名字');
                return;
            }
            if (this.formItem.workplace == "") {
                this.$Message.error('请输入工作地点');
                return;
            }
            this.$http.post('/record/add', this.formItem).then((data) => {
                if (data.data.code == 0) {
                    this.$Message.success("挖掘机出勤记录保存成功");
                    this.$router.push({ path: "/" });
                } else {
                    this.$Message.error("挖掘机出勤记录保存失败");
                }
            });
        }
    }
}
</script>