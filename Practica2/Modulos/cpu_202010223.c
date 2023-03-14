#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/fs.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>

#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/math64.h>
#include <linux/mm.h>
#include <linux/sysinfo.h>


MODULE_LICENSE("GPL"); //LICENCIA GNU GENERAL PUBLIC LICENCE 
MODULE_DESCRIPTION("MODULO DE MEMORIA CPU SO1P2"); // DESCRIPCION 
MODULE_AUTHOR("LUIS ANGEL BARRERA VELASQUEZ 202010223"); // AUTOR 

struct task_struct * cpu;
struct task_struct *procesoHijo;

static int escribir_archivo(struct seq_file *f, void *v){
    unsigned int tiempoTotal = 0;
    unsigned int UsuarioTiempo, SistemaTiempo, UsuarioHijoTiempo, SistemaHijoTiempo;

    rcu_read_lock();

    for_each_process(cpu) {
        if (cpu->mm != NULL) {
            UsuarioTiempo = cpu->utime;
            SistemaTiempo = cpu->stime;
            UsuarioHijoTiempo = cpu->real_parent->utime;
            SistemaHijoTiempo = cpu->real_parent->stime;
            tiempoTotal += UsuarioTiempo + SistemaTiempo + UsuarioHijoTiempo + SistemaHijoTiempo;
        }
    }
    rcu_read_unlock();


    unsigned int porcentajeCPU = ((tiempoTotal*100) / jiffies_64_to_clock_t(HZ))/1000000;
    seq_printf(f, "{\n");
    seq_printf(f, "\"cpu\" : \"%u\",\n", porcentajeCPU);
    seq_printf(f, "\"procesos\" : [");

    unsigned long total_ram;
    struct sysinfo ram;
    si_meminfo(&ram);
    total_ram = ram.totalram * ram.mem_unit;
    
    for_each_process(cpu) {
    int pid = cpu->pid;
    int ppid = cpu->parent->pid;
    // if(cpu->parent->pid!=1){
    //     continue;
    // }
    
    char *nombre = cpu->comm;
    uid_t usuario = cpu->cred->uid.val;
    int estado = cpu->__state;
    int consumoRAM = 0;
    if (cpu->mm) {
        consumoRAM = (cpu->mm->total_vm * PAGE_SIZE) / 1024;
    }

    seq_printf(f, "{");
    seq_printf(f, "\"pid\" : \"%d\",", pid);
    seq_printf(f, "\"ppid\" : \"%d\",", ppid);
    seq_printf(f, "\"nombre\" : \"%s\",", nombre);
    seq_printf(f, "\"usuario\" : \"%d\",", usuario);
    seq_printf(f, "\"estado\" : \"%d\",", estado);
    seq_printf(f, "\"total_ram\" : \"%d\",", total_ram);
    seq_printf(f, "\"ram\" : \"%d\",", consumoRAM);
    seq_printf(f, "\"hijos\" : [");


    list_for_each_entry(procesoHijo, &cpu->children, sibling) {

        int hijoPID = procesoHijo->pid;
        char *hijoNombre = procesoHijo->comm;

        seq_printf(f, "{");
        seq_printf(f, "\"pid\" : \"%d\",", hijoPID);
        seq_printf(f, "\"nombre\" : \"%s\"", hijoNombre);
        seq_printf(f, "}");
    }
    seq_printf(f, "]");


    seq_printf(f, "}");
    }
    seq_printf(f, "]");
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
    proc_create("cpu_202010223", 0, NULL, &operaciones);
    printk(KERN_INFO "LUIS ANGEL BARRERA VELASQUEZ\n");
    return 0;
}

static void remover_modulo(void){
    remove_proc_entry("cpu_202010223", NULL);
    printk(KERN_INFO "Primer Semestre 2023\n");
}

module_init(insertar_modulo);
module_exit(remover_modulo);