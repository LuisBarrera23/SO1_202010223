#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>


#include <linux/hugetlb.h>
#include <linux/mm.h>
#include <linux/mman.h>
#include <linux/mmzone.h>
#include <linux/swap.h>
#include <linux/swapfile.h>
#include <linux/vmstat.h>
#include <linux/atomic.h>
#include <linux/fs.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>



MODULE_LICENSE("GPL"); //LICENCIA GNU GENERAL PUBLIC LICENCE 
MODULE_DESCRIPTION("MODULO DE MEMORIA RAM SO1P2"); // DESCRIPCION 
MODULE_AUTHOR("LUIS ANGEL BARRERA VELASQUEZ 202010223"); // AUTOR 


static int escribir_archivo(struct seq_file *f, void *v){
    struct sysinfo info;
    si_meminfo(&info);
    unsigned long long total = info.totalram;
    unsigned long long libre = info.freeram;
    unsigned long long ocupada = total-libre;
    
    seq_printf(f, "{\n");
    seq_printf(f, "\"memoria_ocupada\": %8llu,\n", ocupada*4096);
    seq_printf(f, "\"memoria_libre\": %8llu,\n", libre*4096);
    seq_printf(f, "\"memoria_total\": %8llu\n", total*4096);
    seq_printf(f, "}\n");
    return 0;
}


//FUNCION QUE SE EJECUTA CADA VEZ QUE SE LEE EL ARCHIVO
static int lectura_llamada(struct inode *inode, struct file *file){
    return single_open(file, escribir_archivo, NULL); 
}

//operaciones que se realizan al leer el archivo, inicializa operaciones
static struct proc_ops operaciones ={
    .proc_open = lectura_llamada,
    .proc_read = seq_read
};

static int insertar_modulo(void){
    proc_create("ram_202010223", 0, NULL, &operaciones);
    printk(KERN_INFO "202010223\n");
    return 0;
}

static void remover_modulo(void){
    remove_proc_entry("ram_202010223", NULL);
    printk(KERN_INFO "SISTEMAS OPERATIVOS 1\n");
}

module_init(insertar_modulo);
module_exit(remover_modulo);