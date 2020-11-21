use futures::Future;
use ocl::{Buffer, Context, Device, Event, Kernel, Platform, ProQue, Program, Queue, RwVec};
use std::fs::File;
use std::io::Write;
use std::thread;

fn boilerplate() -> ocl::Result<()> {
    println!("------ boilerplate ------");
    let platforms = Platform::list();
    for platform in platforms.iter() {
        println!(
            "platform:\n name: {}\n vendor: {}\n",
            platform.name()?,
            platform.vendor()?
        );
        let devices = Device::list_all(platform)?;
        for device in devices {
            println!("device: {}", device.name()?);
        }
    }

    let platform = platforms[0];
    let device = Device::list_all(platform)?[0];
    // let platform = Platform::default();
    // let device = Device::first(platform)?;

    let context = Context::builder()
        .platform(platform)
        .devices(device)
        .build()?;

    let queue = Queue::new(&context, device, None)?;

    let program = Program::builder()
        .src_file("src/add.cl")
        .devices(device)
        .build(&context)?;

    let buffer = Buffer::<f32>::builder()
        .queue(queue.clone())
        .len(5)
        .fill_val(Default::default())
        .build()?;
    buffer.write(&vec![1.0, 2.0, 3.0, 4.0, 5.0]).enq()?;

    let kernel = Kernel::builder()
        .name("add")
        .program(&program)
        .queue(queue.clone())
        .global_work_size(5)
        .arg(&buffer)
        .arg(1.0f32)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut output = vec![0.0f32; 5];
    buffer.read(&mut output).enq()?;

    println!("{:?}", output);
    Ok(())
}

fn trivial() -> ocl::Result<()> {
    println!("------ trivial ------");
    let src = r#"
        __kernel void add(__global float* buffer, float scalar) {
            buffer[get_global_id(0)] += scalar;
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(1 << 10).build()?;

    let buffer = pro_que.create_buffer::<f32>()?;

    let kernel = pro_que
        .kernel_builder("add")
        .arg(&buffer)
        .arg(11.1f32)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut vec = vec![0.0f32; buffer.len()];
    buffer.read(&mut vec).enq()?;

    println!("The value at index [{}] is now '{}'!", 1023, vec[1023]);
    Ok(())
}

fn add_vector() -> ocl::Result<()> {
    println!("------ add vector ------");
    let src = r#"
        __kernel void add_vector(__global float* A, __global float* B,__global float* C) {
            int i=get_global_id(0);
            A[i] = B[i]+C[i];
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(5).build()?;

    let A = pro_que.create_buffer::<f32>()?;
    let B = pro_que.create_buffer::<f32>()?;
    let C = pro_que.create_buffer::<f32>()?;

    B.write(&vec![1.0, 2.0, 3.0, 4.0, 5.0]).enq()?;
    C.write(&vec![1.0, 2.0, 3.0, 4.0, 5.0]).enq()?;

    let kernel = pro_que
        .kernel_builder("add_vector")
        .arg(&A)
        .arg(&B)
        .arg(&C)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut vec = vec![0.0f32; A.len()];
    A.read(&mut vec).enq()?;

    println!("sum is {:?}", vec);
    Ok(())
}

fn global_dimision() -> ocl::Result<()> {
    println!("------ global dimision ------");
    let src = r#"
        __kernel void add(__global float* buffer, float scalar) {
            int id=get_global_id(0)*2;
            buffer[id] += scalar;
            buffer[id+1] += scalar*2;
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(10).build()?;

    let buffer = pro_que.create_buffer::<f32>()?;

    let kernel = pro_que
        .kernel_builder("add")
        .arg(&buffer)
        .arg(1.0f32)
        .global_work_size(5)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut vec = vec![0.0f32; buffer.len()];
    buffer.read(&mut vec).enq()?;

    println!("{:?}", vec);
    Ok(())
}

fn local_dimision_and_barrier() -> ocl::Result<()> {
    println!("------ local dimision ------");
    let src = r#"

__kernel void add(__global float* A, __global float* B) {
    __local float sum[5];
    int gid = get_global_id(0);
    int lid = get_local_id(0);
    sum[lid] = A[gid] + B[gid];

    barrier(CLK_LOCAL_MEM_FENCE);
    if (lid == 0) {
        int i = 0;
        A[gid] = 0;
        for (i = 0; i < 5; i++) {
            A[gid] += sum[i];
        }
    }
}

    "#;

    let pro_que = ProQue::builder().src(src).dims(10).build()?;

    let A = pro_que.create_buffer::<f32>()?;
    let B = pro_que.create_buffer::<f32>()?;

    A.write(&vec![1.0, 2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0, 4.0, 5.0])
        .enq()?;
    B.write(&vec![1.0, 2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0, 4.0, 5.0])
        .enq()?;

    let kernel = pro_que
        .kernel_builder("add")
        .arg(&A)
        .arg(&B)
        .global_work_size(10)
        .local_work_size(5)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut vec = vec![0.0f32; A.len()];
    A.read(&mut vec).enq()?;

    println!("{:?}", vec);
    Ok(())
}

fn reduce_sum() -> ocl::Result<()> {
    println!("------ reduce sum ------");
    let src = r#"

__kernel void reduce_sum(__global float* buffer,__global float* out) {
    int i = get_global_id(0);
    out[i] = buffer[i * 2] + buffer[i * 2 + 1];
}

 "#;

    let pro_que = ProQue::builder().src(src).dims(1024).build()?;
    let mut out = pro_que
        .buffer_builder()
        .len(1024)
        .fill_val(1.0f32)
        .build()?;

    for i in 0..10 {
        let length = 1 << (10 - i - 1);
        let input = out;
        out = pro_que.buffer_builder().len(length).build()?;
        let kernel = pro_que
            .kernel_builder("reduce_sum")
            .arg(&input)
            .arg(&out)
            .global_work_size(length)
            .build()?;

        unsafe {
            kernel.enq()?;
        }
    }

    let mut vec = vec![0.0f32; out.len()];
    out.read(&mut vec).enq()?;

    println!("{:?}", vec[0]);
    Ok(())
}

fn dot() -> ocl::Result<()> {
    println!("------ matrix dot ------");
    let src = r#"

__kernel void dot(__global float* A, __global float* B, __global float* C) {
    int row = get_global_id(0);
    int col = get_global_id(1);
    int i = 0;
    int j = 0;
    for (i = 0; i < 5; i++) {
        C[row * 5 + col] += A[row * 5 + i] * B[i * 5 + col];
    }
}

 "#;

    let pro_que = ProQue::builder().src(src).dims((5, 5)).build()?;

    let A = pro_que.buffer_builder().fill_val(1.0f32).build()?;
    let B = pro_que.buffer_builder().fill_val(1.0f32).build()?;
    let C = pro_que.buffer_builder().fill_val(0.0f32).build()?;

    let kernel = pro_que
        .kernel_builder("dot")
        .arg(&A)
        .arg(&B)
        .arg(&C)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut out = vec![0.0f32; C.len()];
    C.read(&mut out).enq()?;

    println!("{:?}", out);
    Ok(())
}

fn wait_event() -> ocl::Result<()> {
    println!("------ host wait event ------");
    let src = r#"

__kernel void inc(__global float* A) {
    int i = get_global_id(0);
    A[i] += 1.0;
}

 "#;

    let pro_que = ProQue::builder().src(src).dims(10).build()?;
    let A = pro_que.buffer_builder().fill_val(1.0f32).build()?;

    let mut ev = Event::empty();
    let mut out = vec![0.0f32; A.len()];
    let mut rwvec = RwVec::from(out);

    let kernel = pro_que.kernel_builder("inc").arg(&A).build()?;

    unsafe {
        kernel.cmd().enew(&mut ev).enq()?;
    }

    let mut future = A.read(rwvec.clone()).ewait(&ev).enq_async()?;

    thread::spawn(move || {
        future
            .and_then(move |data| {
                println!("{:?}", data.iter().collect::<Vec<&f32>>());
                Ok(())
            })
            .wait();
    })
    .join();

    Ok(())
}

fn mmap() -> ocl::Result<()> {
    println!("------ mmap ------");
    let src = r#"
        __kernel void add(__global float* buffer, float scalar) {
            buffer[get_global_id(0)] += scalar;
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(10).build()?;

    let buffer = pro_que.create_buffer::<f32>()?;

    {
        // mmap for write
        let mut mmap = { unsafe { buffer.map().enq()? } };
        mmap[0] = 1.0f32;
    };

    let kernel = pro_que
        .kernel_builder("add")
        .arg(&buffer)
        .arg(3.14f32)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    // mmap for read
    let mmap = { unsafe { buffer.map().enq()? } };
    println!("{:?}", mmap.iter().cloned().collect::<Vec<f32>>());

    Ok(())
}

use ocl::core;

fn dump_program_binary() -> ocl::Result<()> {
    println!("------ dump program binary ------");
    let src = r#"
        __kernel void add(__global float* buffer, float scalar) {
            buffer[get_global_id(0)] += scalar;
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(10).build()?;

    let binary = core::get_program_info(pro_que.program(), core::ProgramInfo::Binaries)?;
    if let core::ProgramInfoResult::Binaries(data) = binary {
        // device_0 <-> data[0]
        // device_1 <-> data[1]
        println!("{:?}", data.len());
        println!("{:?}", data);
        let mut file = File::create("binary.cl")?;
        file.write_all(data[0].as_slice())?;
    }

    Ok(())
}

fn vector() -> ocl::Result<()> {
    println!("------ trivial ------");
    let src = r#"
        __kernel void add(__global float4* input, float scalar) {
            int i=get_global_id(0);
            input[i].s0 += scalar;
            input[i].s1 += scalar;
            input[i].s2 -= scalar;
            input[i].s3 -= scalar;
        }
    "#;

    let pro_que = ProQue::builder().src(src).dims(1 << 10).build()?;

    let buffer = pro_que
        .buffer_builder()
        .len(4096)
        .copy_host_slice(&vec![0.0f32; 4096])
        .build()?;

    let kernel = pro_que
        .kernel_builder("add")
        .arg(&buffer)
        .arg(10.0f32)
        .build()?;

    unsafe {
        kernel.enq()?;
    }

    let mut vec = vec![0.0f32; buffer.len()];
    buffer.read(&mut vec).enq()?;

    println!("value: {:?}", &vec[0..4]);
    Ok(())
}

// TODO:
// - image & sampler
// - SPIR-V
// - opencl 2.1
//
// install `intel-opencl-icd`, `ocl-icd-opencl`-dev & `clinfo`
// use `clinfo` to make sure opencl is installed correctly
fn main() {
    // boilerplate();
    // trivial();
    // add_vector();
    // global_dimision();
    // local_dimision_and_barrier();
    // reduce_sum()
    // println!("{:?}", dot());
    // wait_event()
    // mmap();
    // dump_program_binary()
    println!("{:?}", vector());
}
