<template>
    <Card>
        <Form>
            <FormItem label="挖掘机" :label-width="80" label-position="top">
                <Select v-model="formItem.device_id">
                    <Option v-for="item in devices" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>

            <FormItem label="时间" :label-width="80" label-position="top">
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
                create_time:"",
            },
            devices: [],
        }
    },
     created() {
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
            if (this.formItem.create_time == "") {
                this.$Message.error('请选择日期时间');
                return;
            }
            this.$http.post('/maintain/add', this.formItem).then((data) => {
                if (data.data.code == 0) {
                    this.$Message.success("挖掘机保养记录保存成功");
                    this.$router.push({ path: "/" });
                } else {
                    this.$Message.error("挖掘机保养记录保存失败");
                }
            });
        }
    }
}
</script>